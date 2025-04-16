package proxy

import (
	"context"
	"github.com/aonescu/glimpse/internal/models"
)

// Proxy defines an interface for any LLM provider
type Proxy interface {
	Call(ctx context.Context, req models.LLMRequest) (models.LLMResponse, error)
}