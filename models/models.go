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
	BodyType      string            `json:"bodyType"` // "json" | "form-data"
}

type RaceResult struct {
	CodeResponse int    `json:"code_response"`
	StatusText   string `json:"status_text"`
	CountCode    int    `json:"count_code"`
	ErrorSample  string `json:"error_sample,omitempty"`
	BodySample   string `json:"body_sample,omitempty"`
}

type RaceSummaryResponse struct {
	TotalRequest int          `json:"total_request"`
	TotalTime    int          `json:"total_time"` // milliseconds
	Result       []RaceResult `json:"result"`
}
