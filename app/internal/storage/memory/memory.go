package memory

import "sync"

// InMemoryStore is a thread-safe map-based store
 type InMemoryStore struct {
	mu sync.RWMutex
	m  map[string]string
 }

 func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{m: make(map[string]string)}
 }

 func (s *InMemoryStore) Save(code, url string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[code] = url
	return nil
 }

 func (s *InMemoryStore) Get(code string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.m[code]
	return v, ok
 }
