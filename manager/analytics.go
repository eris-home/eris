package manager

import (
	"fmt"
	"time"
)

type AnalyticsEvent struct {
	Timestamp time.Time
	EventType string
	Details   map[string]interface{}
}

type AnalyticsManager struct {
	Events []AnalyticsEvent
}

func NewAnalyticsManager() *AnalyticsManager {
	return &AnalyticsManager{
		Events: []AnalyticsEvent{},
	}
}

func (am *AnalyticsManager) RecordEvent(eventType string, details map[string]interface{}) {
	event := AnalyticsEvent{
		Timestamp: time.Now(),
		EventType: eventType,
		Details:   details,
	}
	am.Events = append(am.Events, event)
	fmt.Printf("AnalyticsManager: Event recorded -> %s at %s\n", eventType, event.Timestamp)
}

func (am *AnalyticsManager) GenerateReport() string {
	report := "Analytics Report:\n"
	for _, event := range am.Events {
		report += fmt.Sprintf("[%s] %s: %v\n", event.Timestamp.Format(time.RFC3339), event.EventType, event.Details)
	}
	return report
}

func (am *AnalyticsManager) ClearEvents() {
	am.Events = []AnalyticsEvent{}
	fmt.Println("AnalyticsManager: All events cleared")
}