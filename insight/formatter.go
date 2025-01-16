package insight

import (
	"fmt"
	"strings"
)

type InsightFormatter struct{}

func NewInsightFormatter() *InsightFormatter {
	return &InsightFormatter{}
}

func (f *InsightFormatter) FormatAsText(insights map[string]interface{}) string {
	var builder strings.Builder
	builder.WriteString("Insights Report:\n")
	for key, value := range insights {
		builder.WriteString(fmt.Sprintf("- %s: %v\n", key, value))
	}
	fmt.Println("InsightFormatter: Formatted insights as text")
	return builder.String()
}

func (f *InsightFormatter) FormatAsJSON(insights map[string]interface{}) string {
	json := "{\n"
	for key, value := range insights {
		json += fmt.Sprintf("  \"%s\": \"%v\",\n", key, value)
	}
	json = strings.TrimRight(json, ",\n") + "\n}"
	fmt.Println("InsightFormatter: Formatted insights as JSON")
	return json
}