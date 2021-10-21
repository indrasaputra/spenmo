package model_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/spenmo/internal/repository/model"
)

func TestHookTracing(t *testing.T) {
	query := `DELETE FROM "user_cards" WHERE "user_cards"."id" = \$1`

	t.Run("successfully insert card to database", func(t *testing.T) {
		exec := createCardExecutor()
		defer exec.close()
		exec.client.Use(model.HookTracing)
		exec.mock.ExpectBegin()
		exec.mock.ExpectExec(query).
			WithArgs(sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		exec.mock.ExpectCommit()

		err := exec.client.UserCard.DeleteOneID(1).Exec(testCtx)

		assert.Nil(t, err)
	})
}
