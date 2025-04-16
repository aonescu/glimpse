package main

import (
	"log"
	"net/http"

	"github.com/aonescu/glimpse/internal/api"
	"github.com/aonescu/glimpse/internal/providers"
	"github.com/aonescu/glimpse/internal/proxy"
)

func main() {
	openai := providers.NewOpenAIClient("YOUR_API_KEY")

	dmr := providers.NewDMRClient("http://localhost:12434")

	handler := api.NewHandler(map[string]proxy.Proxy{
	"openai": openai,
	"local":  dmr,
	})

	http.HandleFunc("/v1/llm", handler.HandleLLMProxy)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}