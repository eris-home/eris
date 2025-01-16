package insight

import (
	"fmt"
	"strings"
)

type Analyzer struct{}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) AnalyzeSentiment(input string) string {
	// Mock sentiment analysis
	if strings.Contains(strings.ToLower(input), "happy") {
		return "Positive"
	} else if strings.Contains(strings.ToLower(input), "sad") {
		return "Negative"
	}
	return "Neutral"
}

func (a *Analyzer) AnalyzeTopics(input string) []string {
	// Mock topic detection
	topics := []string{}
	if strings.Contains(strings.ToLower(input), "project") {
		topics = append(topics, "Project Management")
	}
	if strings.Contains(strings.ToLower(input), "deadline") {
		topics = append(topics, "Time Management")
	}
	return topics
}

func (a *Analyzer) AnalyzeComplexity(input string) string {
	// Mock complexity analysis
	wordCount := len(strings.Fields(input))
	switch {
	case wordCount < 10:
		return "Simple"
	case wordCount < 20:
		return "Moderate"
	default:
		return "Complex"
	}
}

func (a *Analyzer) Summarize(input string) string {
	// Mock summary generation
	words := strings.Fields(input)
	if len(words) > 10 {
		return strings.Join(words[:10], " ") + "..."
	}
	return input
}