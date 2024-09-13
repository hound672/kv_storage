package inmemory

import (
	"context"

	"github.com/hound672/kv_storage/internal/models"
)

func (s *InMemoryStorage) Get(ctx context.Context, key string) (string, error) {
	value, ok := s.storage[key]
	if !ok {
		return "", models.ErrKeyNotFound
	}

	return value, nil
}
