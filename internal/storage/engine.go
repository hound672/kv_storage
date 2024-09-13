package storage

import (
	"context"
)

type Storage interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type Engine struct {
	storage Storage
}

func New(storage Storage) *Engine {
	return &Engine{
		storage: storage,
	}
}
