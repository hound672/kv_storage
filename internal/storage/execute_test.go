package storage

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/mock"

	"github.com/hound672/kv_storage/internal/models"
)

func (s *engineSuite) Test_Set_Success() {
	key := gofakeit.Letter()
	value := gofakeit.Letter()

	cmd := models.Command{
		Type: models.CommandTypeSet,
		Args: []string{key, value},
	}

	s.storage.EXPECT().Set(mock.Anything, key, value).Return(nil)

	result, err := s.engine.Execute(cmd)

	s.NoError(err)
	s.Empty(result)
}

func (s *engineSuite) Test_Get_Success() {
	key := gofakeit.Letter()
	value := gofakeit.Letter()

	cmd := models.Command{
		Type: models.CommandTypeGet,
		Args: []string{key},
	}

	s.storage.EXPECT().Get(mock.Anything, key).Return(value, nil)

	result, err := s.engine.Execute(cmd)

	s.NoError(err)
	s.Equal(value, result)
}

func (s *engineSuite) Test_Get_NotFound() {
	key := gofakeit.Letter()

	cmd := models.Command{
		Type: models.CommandTypeGet,
		Args: []string{key},
	}

	s.storage.EXPECT().Get(mock.Anything, key).Return("", models.ErrKeyNotFound)

	result, err := s.engine.Execute(cmd)

	s.ErrorIs(err, models.ErrKeyNotFound)
	s.Empty(result)
}

func (s *engineSuite) Test_Del_Success() {
	key := gofakeit.Letter()

	cmd := models.Command{
		Type: models.CommandTypeDel,
		Args: []string{key},
	}

	s.storage.EXPECT().Delete(mock.Anything, key).Return(nil)

	result, err := s.engine.Execute(cmd)

	s.NoError(err)
	s.Empty(result)
}

func (s *engineSuite) Test_Del_NotFound() {
	key := gofakeit.Letter()

	cmd := models.Command{
		Type: models.CommandTypeDel,
		Args: []string{key},
	}

	s.storage.EXPECT().Delete(mock.Anything, key).Return(models.ErrKeyNotFound)

	result, err := s.engine.Execute(cmd)

	s.ErrorIs(err, models.ErrKeyNotFound)
	s.Empty(result)
}
