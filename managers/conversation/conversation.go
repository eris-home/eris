package conversation

import (
	"fmt"
	"strings"
	"time"
)

type Message struct {
	Timestamp time.Time
	Sender    string
	Content   string
}

type ConversationManager struct {
	Messages    []Message
	ActiveTopic string
}

func NewConversationManager() *ConversationManager {
	return &ConversationManager{
		Messages:    []Message{},
		ActiveTopic: "",
	}
}

func (cm *ConversationManager) AddMessage(sender, content string) {
	message := Message{
		Timestamp: time.Now(),
		Sender:    sender,
		Content:   content,
	}
	cm.Messages = append(cm.Messages, message)
	fmt.Printf("ConversationManager: Message added -> [%s] %s: %s\n", message.Timestamp.Format(time.RFC3339), sender, content)
}

func (cm *ConversationManager) ListMessages() {
	fmt.Println("ConversationManager: Listing all messages")
	for _, msg := range cm.Messages {
		fmt.Printf("[%s] %s: %s\n", msg.Timestamp.Format(time.RFC3339), msg.Sender, msg.Content)
	}
}

func (cm *ConversationManager) DetectTopic(input string) string {
	lowerInput := strings.ToLower(input)
	if strings.Contains(lowerInput, "project") {
		cm.ActiveTopic = "Project Discussion"
	} else if strings.Contains(lowerInput, "deadline") {
		cm.ActiveTopic = "Deadline Management"
	} else if strings.Contains(lowerInput, "help") {
		cm.ActiveTopic = "Help and Support"
	} else {
		cm.ActiveTopic = "General Discussion"
	}
	fmt.Printf("ConversationManager: Active topic detected -> %s\n", cm.ActiveTopic)
	return cm.ActiveTopic
}

func (cm *ConversationManager) SummarizeConversation() string {
	var summary strings.Builder
	summary.WriteString("Conversation Summary:\n")

	for _, msg := range cm.Messages {
		summary.WriteString(fmt.Sprintf("[%s] %s: %s\n", msg.Timestamp.Format("15:04"), msg.Sender, msg.Content))
	}

	fmt.Println("ConversationManager: Conversation summarized")
	return summary.String()
}

func (cm *ConversationManager) ClearConversation() {
	cm.Messages = []Message{}
	cm.ActiveTopic = ""
	fmt.Println("ConversationManager: Conversation cleared")
}