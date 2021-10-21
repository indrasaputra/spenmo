package model_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/spenmo/internal/repository/model"
)

func TestHookTracing(t *testing.T) {
	query := `INSERT INTO "user_cards" \("created_at", "updated_at", "limit_daily", "limit_monthly", "user_id", "wallet_id"\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6\) RETURNING "id"`

	t.Run("successfully insert card to database", func(t *testing.T) {
		exec := createCardExecutor()
		defer exec.close()
		exec.client.Use(model.HookTracing)
		exec.mock.ExpectBegin()
		exec.mock.ExpectQuery(query).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		exec.mock.ExpectCommit()
		card := model.NewCard(exec.client)

		err := card.Insert(testCtx, createCard())

		assert.Nil(t, err)
	})
}
