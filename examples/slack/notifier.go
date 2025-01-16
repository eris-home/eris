package slack

import "fmt"

type Notifier struct {
	Logger *SlackLogger
}

func NewNotifier(logger *SlackLogger) *Notifier {
	return &Notifier{
		Logger: logger,
	}
}

func (n *Notifier) SendNotification(message string) {
	fmt.Printf("\033[33m[NOTIFICATION]\033[0m %s\n", message)
	n.Logger.LogInfo("Notifier: " + message)
}