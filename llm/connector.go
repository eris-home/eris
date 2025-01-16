package llm

import (
	"fmt"
	"net/http"
	"time"
)

type LLMConnector struct {
	BaseURL string
	APIKey  string
	Client  *http.Client
}

func NewLLMConnector(baseURL, apiKey string) *LLMConnector {
	return &LLMConnector{
		BaseURL: baseURL,
		APIKey:  apiKey,
		Client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (c *LLMConnector) SendRequest(endpoint string, payload map[string]interface{}) (string, error) {
	// Simulate sending a request to the LLM API
	url := fmt.Sprintf("%s/%s", c.BaseURL, endpoint)
	fmt.Printf("LLMConnector: Sending request to %s with payload %v\n", url, payload)

	// Mock response
	response := fmt.Sprintf("Response from %s", url)
	return response, nil
}

func (c *LLMConnector) CheckHealth() bool {
	// Mock health check
	fmt.Printf("LLMConnector: Checking health of %s\n", c.BaseURL)
	return true
}