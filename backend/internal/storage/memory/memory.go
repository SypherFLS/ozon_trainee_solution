package memory

import (
	"errors"
	"sync"
)

type Storage struct {
    mu sync.RWMutex

    byOriginal map[string]string
    byShort    map[string]string
}

func New() *Storage {
    return &Storage{
        byOriginal: make(map[string]string),
        byShort: make(map[string]string),
    }
}

func (s *Storage) Save(original, short string) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.byOriginal[original] = short
    s.byShort[short] = original

    return nil
}

func (s *Storage) GetByOriginal(original string) (string, error) {
	s.mu.RLock()
    defer s.mu.RUnlock()

	short, ok := s.byOriginal[original] 
	if !ok {
		return "", errors.New("not found")
	}

	return short, nil
}

func (s *Storage) GetByShort(short string) (string, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    original, ok := s.byShort[short]
    if !ok {
        return "", errors.New("not found")
    }

    return original, nil
}

func (s *Storage) ExistsShort(short string) bool {
    s.mu.RLock()
    defer s.mu.RUnlock()
    _, ok := s.byShort[short]
    return ok
}