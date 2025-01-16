package engine

import (
	"fmt"
	"strings"
)

type Utils struct{}

func NewUtils() *Utils {
	return &Utils{}
}

func (u *Utils) NormalizeInput(input string) string {
	normalized := strings.TrimSpace(strings.ToLower(input))
	fmt.Printf("Utils: Normalized input '%s' to '%s'\n", input, normalized)
	return normalized
}

func (u *Utils) ContainsKeyword(input, keyword string) bool {
	contains := strings.Contains(strings.ToLower(input), strings.ToLower(keyword))
	fmt.Printf("Utils: Input '%s' contains keyword '%s': %t\n", input, keyword, contains)
	return contains
}