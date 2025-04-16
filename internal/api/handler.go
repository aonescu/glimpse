package api

import (
	"encoding/json"
	"net/http"

	"github.com/aonescu/glimpse/internal/models"
	"github.com/aonescu/glimpse/internal/proxy"
)

type Handler struct {
	Providers map[string]proxy.Proxy
}

func NewHandler(providers map[string]proxy.Proxy) *Handler {
	return &Handler{Providers: providers}
}

func (h *Handler) HandleLLMProxy(w http.ResponseWriter, r *http.Request) {
	var req models.LLMRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	provider, ok := h.Providers[req.Provider]
	if !ok {
		http.Error(w, "Unsupported provider", http.StatusBadRequest)
		return
	}

	resp, err := provider.Call(r.Context(), req)
	if err != nil {
		http.Error(w, "LLM error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}