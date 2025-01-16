package state

import (
	"fmt"
	"time"
)

type Session struct {
	ID        string
	ActorID   string
	CreatedAt time.Time
	LastSeen  time.Time
	Data      map[string]interface{}
}

func NewSession(id, actorID string) *Session {
	now := time.Now()
	return &Session{
		ID:        id,
		ActorID:   actorID,
		CreatedAt: now,
		LastSeen:  now,
		Data:      make(map[string]interface{}),
	}
}

func (s *Session) UpdateLastSeen() {
	s.LastSeen = time.Now()
	fmt.Printf("Session %s last seen updated to %s\n", s.ID, s.LastSeen)
}

func (s *Session) SetData(key string, value interface{}) {
	s.Data[key] = value
	fmt.Printf("Session data set: %s = %v\n", key, value)
}

func (s *Session) GetData(key string) (interface{}, bool) {
	value, exists := s.Data[key]
	return value, exists
}

func (s *Session) ResetData() {
	s.Data = make(map[string]interface{})
	fmt.Printf("Session data reset for ID: %s\n", s.ID)
}