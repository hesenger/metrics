package auth

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"
)

type stateEntry struct {
	createdAt time.Time
}

type StateStore struct {
	mu     sync.RWMutex
	states map[string]stateEntry
}

func NewStateStore() *StateStore {
	store := &StateStore{
		states: make(map[string]stateEntry),
	}
	go store.cleanup()
	return store
}

func (s *StateStore) GenerateState() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	state := base64.URLEncoding.EncodeToString(b)

	s.mu.Lock()
	s.states[state] = stateEntry{createdAt: time.Now()}
	s.mu.Unlock()

	return state, nil
}

func (s *StateStore) ValidateState(state string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	entry, exists := s.states[state]
	if !exists {
		return false
	}

	if time.Since(entry.createdAt) > 5*time.Minute {
		delete(s.states, state)
		return false
	}

	delete(s.states, state)
	return true
}

func (s *StateStore) cleanup() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		s.mu.Lock()
		for state, entry := range s.states {
			if time.Since(entry.createdAt) > 5*time.Minute {
				delete(s.states, state)
			}
		}
		s.mu.Unlock()
	}
}
