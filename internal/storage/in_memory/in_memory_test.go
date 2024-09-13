package inmemory

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type inMemorySuite struct {
	suite.Suite

	storage *InMemoryStorage
}

func (s *inMemorySuite) SetupTest() {
	s.storage = New()
}

func TestRunInMemorySuite(t *testing.T) {
	suite.Run(t, new(inMemorySuite))
}
