package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aonescu/glimpse/internal/models"
)

type DockerModelRunnerClient struct {
	BaseURL string // e.g., http://localhost:12434
	Client  *http.Client
}

func NewDMRClient(baseURL string) *DockerModelRunnerClient {
	return &DockerModelRunnerClient{
		BaseURL: baseURL,
		Client:  &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *DockerModelRunnerClient) Call(ctx context.Context, req models.LLMRequest) (models.LLMResponse, error) {
	start := time.Now()

	payload := map[string]interface{}{
		"model": req.Model,
		"messages": []map[string]string{
			{"role": "system", "content": "You are a helpful assistant."},
			{"role": "user", "content": req.Prompt},
		},
	}

	body, _ := json.Marshal(payload)
	url := fmt.Sprintf("%s/engines/llama.cpp/v1/chat/completions", c.BaseURL)

	httpReq, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(httpReq)
	if err != nil {
		return models.LLMResponse{}, err
	}
	defer resp.Body.Close()

	var data struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
		Usage struct {
			TotalTokens int `json:"total_tokens"`
		} `json:"usage"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return models.LLMResponse{}, err
	}

	return models.LLMResponse{
		Output:       data.Choices[0].Message.Content,
		TokenUsage:   data.Usage.TotalTokens,
		Cost:         0, // Local inference = free
		ResponseTime: time.Since(start).Milliseconds(),
	}, nil
}