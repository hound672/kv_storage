package inmemory

import (
	"context"

	"github.com/hound672/kv_storage/internal/models"
)

func (s *InMemoryStorage) Delete(_ context.Context, key string) error {
	if _, ok := s.storage[key]; !ok {
		return models.ErrKeyNotFound
	}

	delete(s.storage, key)
	return nil
}
