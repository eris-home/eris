package main

import (
	"github.com/eris-home/eris/examples/slack"
)

func main() {
	// Slack bot setup
	token := "xoxb-your-slack-token"
	listener := slack.NewSlackListener(token)
	logger := slack.NewSlackLogger()
	responder := slack.NewResponder()
	notifier := slack.NewNotifier(logger)

	logger.LogInfo("Starting Slack bot...")
	listener.StartListening()

	// Example response
	responder.RespondToMessage("Hello, bot!", func(response string) {
		logger.LogInfo("Responder: Sending response -> " + response)
	})

	// Example notification
	notifier.SendNotification("New message received!")

	// Prevent exit
	select {}
}