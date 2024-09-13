package inmemory

import (
	"context"
)

func (s *InMemoryStorage) Set(_ context.Context, key, value string) error {
	s.storage[key] = value
	return nil
}
