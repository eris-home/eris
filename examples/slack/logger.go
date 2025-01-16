package slack

import (
	"fmt"
	"log"
	"os"
)

type SlackLogger struct {
	Logger *log.Logger
}

func NewSlackLogger() *SlackLogger {
	file, err := os.OpenFile("slack_bot.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	logger := log.New(file, "SLACK-BOT: ", log.Ldate|log.Ltime|log.Lshortfile)
	return &SlackLogger{
		Logger: logger,
	}
}

func (sl *SlackLogger) LogInfo(message string) {
	fmt.Printf("\033[34m[INFO]\033[0m %s\n", message)
	sl.Logger.Println("[INFO] " + message)
}

func (sl *SlackLogger) LogError(message string) {
	fmt.Printf("\033[31m[ERROR]\033[0m %s\n", message)
	sl.Logger.Println("[ERROR] " + message)
}

func (sl *SlackLogger) LogDebug(message string) {
	fmt.Printf("\033[32m[DEBUG]\033[0m %s\n", message)
	sl.Logger.Println("[DEBUG] " + message)
}