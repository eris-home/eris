package manager

import (
	"fmt"
	"sync"
)

type Event struct {
	Name    string
	Payload map[string]interface{}
}

type Observer interface {
	Notify(event Event)
}

type EventManager struct {
	subscribers []Observer
	mu          sync.RWMutex
}

func NewEventManager() *EventManager {
	return &EventManager{subscribers: []Observer{}}
}

func (em *EventManager) Subscribe(observer Observer) {
	em.mu.Lock()
	defer em.mu.Unlock()

	em.subscribers = append(em.subscribers, observer)
	fmt.Println("EventManager: Observer subscribed")
}

func (em *EventManager) Unsubscribe(observer Observer) {
	em.mu.Lock()
	defer em.mu.Unlock()

	for i, sub := range em.subscribers {
		if sub == observer {
			em.subscribers = append(em.subscribers[:i], em.subscribers[i+1:]...)
			fmt.Println("EventManager: Observer unsubscribed")
			break
		}
	}
}

func (em *EventManager) NotifyAll(event Event) {
	em.mu.RLock()
	defer em.mu.RUnlock()

	fmt.Printf("EventManager: Notifying observers of event '%s'\n", event.Name)
	for _, observer := range em.subscribers {
		observer.Notify(event)
	}
}