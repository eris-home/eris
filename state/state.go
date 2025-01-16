package state

import (
	"fmt"
	"sync"
)

type StateManager struct {
	GlobalState map[string]interface{}
	mu          sync.RWMutex
}

func NewStateManager() *StateManager {
	return &StateManager{
		GlobalState: make(map[string]interface{}),
	}
}

func (sm *StateManager) Set(key string, value interface{}) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.GlobalState[key] = value
	fmt.Printf("Global state set: %s = %v\n", key, value)
}

func (sm *StateManager) Get(key string) (interface{}, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	value, exists := sm.GlobalState[key]
	return value, exists
}

func (sm *StateManager) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	delete(sm.GlobalState, key)
	fmt.Printf("Global state deleted: %s\n", key)
}

func (sm *StateManager) ListKeys() []string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	var keys []string
	for key := range sm.GlobalState {
		keys = append(keys, key)
	}
	return keys
}