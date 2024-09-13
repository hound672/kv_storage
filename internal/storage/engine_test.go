package storage

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/hound672/kv_storage/internal/storage/mocks"
)

type engineSuite struct {
	suite.Suite

	storage *mocks.Storage

	engine *Engine
}

func (s *engineSuite) SetupTest() {
	t := s.T()

	s.storage = mocks.NewStorage(t)

	s.engine = New(s.storage)
}

func TestRunEngineSuite(t *testing.T) {
	suite.Run(t, new(engineSuite))
}
