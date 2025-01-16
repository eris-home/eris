package insight

import (
	"fmt"
	"time"
)

type Insight struct {
	ID        string
	Timestamp time.Time
	Content   string
	Category  string
}

type InsightManager struct {
	Insights []Insight
}

func NewInsightManager() *InsightManager {
	return &InsightManager{
		Insights: []Insight{},
	}
}

func (im *InsightManager) AddInsight(content, category string) {
	insight := Insight{
		ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
		Timestamp: time.Now(),
		Content:   content,
		Category:  category,
	}
	im.Insights = append(im.Insights, insight)
	fmt.Printf("InsightManager: Added new insight [%s] in category '%s'\n", insight.ID, category)
}

func (im *InsightManager) ListInsights() {
	fmt.Println("InsightManager: Listing all insights")
	for _, insight := range im.Insights {
		fmt.Printf("[%s] %s: %s\n", insight.Timestamp.Format(time.RFC3339), insight.Category, insight.Content)
	}
}

func (im *InsightManager) ClearInsights() {
	im.Insights = []Insight{}
	fmt.Println("InsightManager: Cleared all insights")
}