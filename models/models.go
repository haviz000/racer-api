package models

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RaceRequest struct {
	URL           string            `json:"url"`
	Method        string            `json:"method"`
	Total         int               `json:"total_request"`
	Concurrent    int               `json:"concurrent"`
	Payload       map[string]any    `json:"payload"`
	Headers       map[string]string `json:"headers"`
	Authorization string            `json:"authorization"`
}

type RaceResult struct {
	CodeResponse int `json:"code_response"`
	CountCode    int `json:"count_code"`
}

type RaceSummaryResponse struct {
	TotalRequest int          `json:"total_request"`
	TotalTime    int          `json:"total_time"` // milliseconds
	Result       []RaceResult `json:"result"`
}
