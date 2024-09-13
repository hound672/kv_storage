package models

import (
	"errors"
)

var (
	ErrInvalidCommand = errors.New("invalid command")
	ErrKeyNotFound    = errors.New("key not found")
)
