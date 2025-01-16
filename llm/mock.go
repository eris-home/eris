package llm

import "fmt"

type MockProvider struct {
	Name string
}

func NewMockProvider(name string) *MockProvider {
	return &MockProvider{Name: name}
}

func (m *MockProvider) GenerateText(request LLMRequest) (LLMResponse, error) {
	// Simulate text generation
	response := LLMResponse{
		Text:       fmt.Sprintf("Mock response for prompt: '%s'", request.Prompt),
		TokensUsed: CountTokens(request.Prompt),
	}
	fmt.Printf("MockProvider: Generated response with %d tokens\n", response.TokensUsed)
	return response, nil
}

func (m *MockProvider) GetName() string {
	return m.Name
}