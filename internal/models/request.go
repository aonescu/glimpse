package models

type LLMRequest struct {
	UserID     string
	Provider   string
	Model      string
	Prompt     string
	Temperature float32
	MaxTokens  int
}

type LLMResponse struct {
	Output       string
	TokenUsage   int
	Cost         float64
	ResponseTime int64 // ms
}