package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/spenmo/entity"
)

func TestErrInternal(t *testing.T) {
	t.Run("success get internal error", func(t *testing.T) {
		err := entity.ErrInternal("")
		assert.Contains(t, err.Error(), "rpc error: code = Internal")
	})
}

func TestErrEmptyCard(t *testing.T) {
	t.Run("success get empty card error", func(t *testing.T) {
		err := entity.ErrEmptyCard()
		assert.Contains(t, err.Error(), "rpc error: code = InvalidArgument")
	})
}

func TestErrInvalidID(t *testing.T) {
	t.Run("success get invalid id error", func(t *testing.T) {
		err := entity.ErrInvalidID()
		assert.Contains(t, err.Error(), "rpc error: code = InvalidArgument")
	})
}

func TestErrNotFound(t *testing.T) {
	t.Run("success get not found error", func(t *testing.T) {
		err := entity.ErrNotFound()
		assert.Contains(t, err.Error(), "rpc error: code = NotFound")
	})
}
