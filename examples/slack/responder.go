package slack

import (
	"fmt"
	"strings"
)

type Responder struct{}

func NewResponder() *Responder {
	return &Responder{}
}

func (r *Responder) RespondToMessage(message string, callback func(response string)) {
	fmt.Printf("Responder: Processing message -> %s\n", message)

	// Simple mock response logic
	if strings.Contains(strings.ToLower(message), "hello") {
		callback("Hi there! How can I assist you today?")
	} else {
		callback("I'm sorry, I didn't understand that. Could you clarify?")
	}
}