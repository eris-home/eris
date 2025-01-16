package state

import "fmt"

type Event struct {
	Type      string
	Timestamp string
	Data      map[string]interface{}
}

type Tracker struct {
	Events []Event
}

func NewTracker() *Tracker {
	return &Tracker{Events: []Event{}}
}

func (t *Tracker) RecordEvent(eventType, timestamp string, data map[string]interface{}) {
	event := Event{
		Type:      eventType,
		Timestamp: timestamp,
		Data:      data,
	}
	t.Events = append(t.Events, event)
	fmt.Printf("Event recorded: %s at %s\n", eventType, timestamp)
}

func (t *Tracker) ListEvents() []Event {
	fmt.Println("Listing all events:")
	for _, event := range t.Events {
		fmt.Printf("Type: %s, Timestamp: %s, Data: %v\n", event.Type, event.Timestamp, event.Data)
	}
	return t.Events
}

func (t *Tracker) ClearEvents() {
	t.Events = []Event{}
	fmt.Println("All events cleared")
}