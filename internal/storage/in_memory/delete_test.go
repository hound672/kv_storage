package inmemory

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/hound672/kv_storage/internal/models"
)

func (s *inMemorySuite) Test_Delete() {
	key := gofakeit.Word()

	s.storage.storage[key] = gofakeit.Word()

	err := s.storage.Delete(context.Background(), key)
	s.NoError(err)

	_, ok := s.storage.storage[key]
	s.False(ok)
}

func (s *inMemorySuite) Test_Delete_NotFound() {
	err := s.storage.Delete(context.Background(), gofakeit.Word())

	s.ErrorIs(err, models.ErrKeyNotFound)
}
