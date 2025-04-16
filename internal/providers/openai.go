package providers

import (
	"context"
	"time"

	"github.com/aonescu/glimpse/internal/models"
)

type OpenAIClient struct {
	APIKey string
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{APIKey: apiKey}
}

func (c *OpenAIClient) Call(ctx context.Context, req models.LLMRequest) (models.LLMResponse, error) {
	// TODO: Real HTTP call to OpenAI's endpoint
	start := time.Now()

	resp := models.LLMResponse{
		Output:       "This is a mock response from OpenAI",
		TokenUsage:   100,
		Cost:         0.002,
		ResponseTime: time.Since(start).Milliseconds(),
	}

	return resp, nil
}