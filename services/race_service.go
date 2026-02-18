package services

import (
	"bytes"
	"encoding/json"
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
	var mu sync.Mutex

	client := &http.Client{}
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

			mu.Lock()
			resultMap[resp.StatusCode]++
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	totalTime := int(time.Since(start).Milliseconds())

	results := make([]models.RaceResult, 0, len(resultMap))
	for code, count := range resultMap {
		results = append(results, models.RaceResult{
			CodeResponse: code,
			CountCode:    count,
		})
	}

	return models.RaceSummaryResponse{
		TotalRequest: req.Total,
		TotalTime:    totalTime,
		Result:       results,
	}
}
