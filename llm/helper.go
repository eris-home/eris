package llm

import (
	"fmt"
	"strings"
)

func CountTokens(input string) int {
	// Simulate token counting
	tokens := strings.Fields(input)
	tokenCount := len(tokens)
	fmt.Printf("Helper: Counted %d tokens in input\n", tokenCount)
	return tokenCount
}

func SplitIntoChunks(input string, maxTokens int) []string {
	// Split input into chunks based on maxTokens
	words := strings.Fields(input)
	var chunks []string
	for len(words) > 0 {
		if len(words) <= maxTokens {
			chunks = append(chunks, strings.Join(words, " "))
			break
		}
		chunks = append(chunks, strings.Join(words[:maxTokens], " "))
		words = words[maxTokens:]
	}
	fmt.Printf("Helper: Split input into %d chunks\n", len(chunks))
	return chunks
}