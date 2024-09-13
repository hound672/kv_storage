package inmemory

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/hound672/kv_storage/internal/models"
)

func (s *inMemorySuite) Test_Get() {
	key := gofakeit.Word()
	value := gofakeit.Word()

	s.storage.storage[key] = value

	result, err := s.storage.Get(context.Background(), key)

	s.NoError(err)
	s.Equal(value, result)
}

func (s *inMemorySuite) Test_Get_NotFound() {
	result, err := s.storage.Get(context.Background(), gofakeit.Word())

	s.ErrorIs(err, models.ErrKeyNotFound)
	s.Empty(result)
}
