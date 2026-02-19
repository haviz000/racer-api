package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/haviz000/racer-api/models"
)

func ExecuteRaceTest(req models.RaceRequest) models.RaceSummaryResponse {
	if req.Concurrent <= 0 {
		req.Concurrent = 1
	}

	wg := sync.WaitGroup{}
	sem := make(chan struct{}, req.Concurrent)

	resultMap := make(map[int]int)
	errorMap := make(map[int]string)
	bodyMap := make(map[int]string)
	var mu sync.Mutex

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	start := time.Now()

	for i := 0; i < req.Total; i++ {
		wg.Add(1)

		sem <- struct{}{} // aman sekarang

		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }()

			body, _ := json.Marshal(req.Payload)
			httpReq, err := http.NewRequest(req.Method, req.URL, bytes.NewBuffer(body))
			if err != nil {
				mu.Lock()
				resultMap[0]++
				if _, exists := errorMap[0]; !exists {
					errorMap[0] = err.Error()
				}
				mu.Unlock()
				return
			}

			for k, v := range req.Headers {
				httpReq.Header.Set(k, v)
			}

			if req.Authorization != "" {
				httpReq.Header.Set("Authorization", req.Authorization)
			}

			resp, err := client.Do(httpReq)
			if err != nil {
				mu.Lock()
				resultMap[0]++
				mu.Unlock()
				return
			}
			defer resp.Body.Close()

			bodyBytes, _ := io.ReadAll(resp.Body)
			bodyStr := string(bodyBytes)

			mu.Lock()
			resultMap[resp.StatusCode]++
			if _, exists := bodyMap[resp.StatusCode]; !exists {
				if len(bodyStr) > 1000 {
					bodyStr = bodyStr[:1000] + "...(truncated)"
				}
				bodyMap[resp.StatusCode] = bodyStr
			}
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	totalTime := int(time.Since(start).Milliseconds())

	results := make([]models.RaceResult, 0, len(resultMap))
	for code, count := range resultMap {
		statusText := http.StatusText(code)

		if code == 0 {
			statusText = "Network / Request Error"
		}

		if statusText == "" {
			statusText = "Unknown Status"
		}

		results = append(results, models.RaceResult{
			CodeResponse: code,
			StatusText:   statusText,
			CountCode:    count,
			ErrorSample:  errorMap[code],
			BodySample:   bodyMap[code],
		})
	}

	return models.RaceSummaryResponse{
		TotalRequest: req.Total,
		TotalTime:    totalTime,
		Result:       results,
	}
}
