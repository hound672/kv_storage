package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hound672/kv_storage/internal/models"
)

func Test_Parse(t *testing.T) {
	tests := []struct {
		name        string
		source      string
		expected    models.Command
		expectedErr error
	}{
		// SET
		{
			name:   "set success",
			source: "SET key value",
			expected: models.Command{
				Type: models.CommandTypeSet,
				Args: []string{"key", "value"},
			},
			expectedErr: nil,
		},
		{
			name:        "set invalid. No args",
			source:      "SET",
			expected:    models.Command{},
			expectedErr: models.ErrInvalidCommand,
		},
		{
			name:        "set invalid. Not enough args",
			source:      "SET key",
			expected:    models.Command{},
			expectedErr: models.ErrInvalidCommand,
		},
		{
			name:        "set invalid. Too many args",
			source:      "SET key1 key2 key3",
			expected:    models.Command{},
			expectedErr: models.ErrInvalidCommand,
		},
		// GET
		{
			name:   "get success",
			source: "GET key",
			expected: models.Command{
				Type: models.CommandTypeGet,
				Args: []string{"key"},
			},
			expectedErr: nil,
		},
		{
			name:        "get invalid. No args",
			source:      "GET",
			expected:    models.Command{},
			expectedErr: models.ErrInvalidCommand,
		},
		{
			name:        "get invalid. Too many args.",
			source:      "GET key1 key2",
			expected:    models.Command{},
			expectedErr: models.ErrInvalidCommand,
		},
		// DEL
		{
			name:   "del success",
			source: "DEL key",
			expected: models.Command{
				Type: models.CommandTypeDel,
				Args: []string{"key"},
			},
			expectedErr: nil,
		},
		{
			name:        "del invalid. No args",
			source:      "DEL",
			expected:    models.Command{},
			expectedErr: models.ErrInvalidCommand,
		},
		{
			name:        "del invalid. Too many args.",
			source:      "DEL key1 key2",
			expected:    models.Command{},
			expectedErr: models.ErrInvalidCommand,
		},
	}

	parser := New()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := parser.Parse(tc.source)
			if tc.expectedErr != nil {
				require.ErrorIs(t, err, tc.expectedErr)
				return
			}

			require.Equal(t, tc.expected, result)
		})
	}
}
