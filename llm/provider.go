package llm

import (
	"fmt"
)

type Provider struct {
	Name      string
	Connector *LLMConnector
}

func NewProvider(name, baseURL, apiKey string) *Provider {
	connector := NewLLMConnector(baseURL, apiKey)
	return &Provider{
		Name:      name,
		Connector: connector,
	}
}

func (p *Provider) GenerateText(request LLMRequest) (LLMResponse, error) {
	payload := map[string]interface{}{
		"prompt":     request.Prompt,
		"max_tokens": request.MaxTokens,
	}
	response, err := p.Connector.SendRequest("generate", payload)
	if err != nil {
		return LLMResponse{}, fmt.Errorf("failed to generate text: %w", err)
	}

	// Simulate parsing response
	tokensUsed := CountTokens(request.Prompt)
	fmt.Printf("Provider '%s': Text generated with %d tokens\n", p.Name, tokensUsed)
	return LLMResponse{Text: response, TokensUsed: tokensUsed}, nil
}

func (p *Provider) GetName() string {
	return p.Name
}