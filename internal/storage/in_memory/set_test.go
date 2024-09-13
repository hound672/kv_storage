package inmemory

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
)

func (s *inMemorySuite) Test_Set() {
	key := gofakeit.Word()
	value := gofakeit.Word()

	err := s.storage.Set(context.Background(), key, value)
	s.NoError(err)

	s.Equal(value, s.storage.storage[key])
}
