package slack

import (
	"fmt"
	"log"
)

type SlackListener struct {
	Token string
}

func NewSlackListener(token string) *SlackListener {
	return &SlackListener{
		Token: token,
	}
}

func (sl *SlackListener) StartListening() {
	fmt.Println("SlackListener: Listening for Slack events...")
	// Mock event listening
	go func() {
		// Simulate listening to Slack events
		for {
			event := "Mock Slack Event: User Message"
			fmt.Printf("SlackListener: Received event -> %s\n", event)
		}
	}()
}

func (sl *SlackListener) StopListening() {
	fmt.Println("SlackListener: Stopped listening for Slack events.")
}