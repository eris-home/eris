package llm

import "fmt"

type LLMRequest struct {
	Prompt string
	MaxTokens int
}

type LLMResponse struct {
	Text string
	TokensUsed int
}

type LLM interface {
	GenerateText(request LLMRequest) (LLMResponse, error)
	GetName() string
}

type LLMManager struct {
	Providers map[string]LLM
}

func NewLLMManager() *LLMManager {
	return &LLMManager{
		Providers: make(map[string]LLM),
	}
}

func (m *LLMManager) RegisterProvider(name string, provider LLM) {
	m.Providers[name] = provider
	fmt.Printf("LLMManager: Registered provider '%s'\n", name)
}

func (m *LLMManager) GenerateWithProvider(name string, request LLMRequest) (LLMResponse, error) {
	provider, exists := m.Providers[name]
	if !exists {
		return LLMResponse{}, fmt.Errorf("provider '%s' not found", name)
	}
	fmt.Printf("LLMManager: Generating text with provider '%s'\n", name)
	return provider.GenerateText(request)
}