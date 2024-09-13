package storage

import (
	"context"
	"fmt"

	"github.com/hound672/kv_storage/internal/models"
)

func (e *Engine) Execute(cmd models.Command) (string, error) {
	ctx := context.Background()

	switch cmd.Type {
	case models.CommandTypeSet:
		if err := e.storage.Set(ctx, cmd.Args[0], cmd.Args[1]); err != nil {
			return "", fmt.Errorf("e.storage.Set: %w", err)
		}
		return "", nil
	case models.CommandTypeGet:
		value, err := e.storage.Get(ctx, cmd.Args[0])
		if err != nil {
			return "", fmt.Errorf("e.storage.Get: %w", err)
		}
		return value, nil
	case models.CommandTypeDel:
		if err := e.storage.Delete(ctx, cmd.Args[0]); err != nil {
			return "", fmt.Errorf("e.storage.Delete: %w", err)
		}
		return "", nil
	}

	return "", nil
}
