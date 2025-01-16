package engine

import (
	"fmt"
	"time"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(message string) {
	fmt.Printf("[%s] INFO: %s\n", time.Now().Format(time.RFC3339), message)
}

func (l *Logger) Debug(message string) {
	fmt.Printf("[%s] DEBUG: %s\n", time.Now().Format(time.RFC3339), message)
}

func (l *Logger) Error(message string) {
	fmt.Printf("[%s] ERROR: %s\n", time.Now().Format(time.RFC3339), message)
}