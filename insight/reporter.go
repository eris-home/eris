package insight

import (
	"fmt"
)

type Reporter struct {
	Formatter *InsightFormatter
}

func NewReporter(formatter *InsightFormatter) *Reporter {
	return &Reporter{Formatter: formatter}
}

func (r *Reporter) GenerateReport(insights []Insight) string {
	insightsMap := make(map[string]interface{})
	for _, insight := range insights {
		insightsMap[insight.Category] = append(insightsMap[insight.Category].([]string), insight.Content)
	}

	report := r.Formatter.FormatAsText(insightsMap)
	fmt.Println("Reporter: Generated insights report")
	return report
}

func (r *Reporter) GenerateJSONReport(insights []Insight) string {
	insightsMap := make(map[string]interface{})
	for _, insight := range insights {
		insightsMap[insight.Category] = append(insightsMap[insight.Category].([]string), insight.Content)
	}

	jsonReport := r.Formatter.FormatAsJSON(insightsMap)
	fmt.Println("Reporter: Generated JSON report")
	return jsonReport
}