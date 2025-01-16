package insight

import (
	"fmt"
)

type InsightExtractor struct{}

func NewInsightExtractor() *InsightExtractor {
	return &InsightExtractor{}
}

func (ie *InsightExtractor) ExtractSentiment(input string) string {
	analyzer := NewAnalyzer()
	sentiment := analyzer.AnalyzeSentiment(input)
	fmt.Printf("InsightExtractor: Extracted sentiment '%s' from input\n", sentiment)
	return sentiment
}

func (ie *InsightExtractor) ExtractTopics(input string) []string {
	analyzer := NewAnalyzer()
	topics := analyzer.AnalyzeTopics(input)
	fmt.Printf("InsightExtractor: Extracted topics '%v' from input\n", topics)
	return topics
}

func (ie *InsightExtractor) ExtractSummary(input string) string {
	analyzer := NewAnalyzer()
	summary := analyzer.Summarize(input)
	fmt.Printf("InsightExtractor: Extracted summary '%s'\n", summary)
	return summary
}