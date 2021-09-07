package postgres_test

import (
	"context"
	"errors"
	"log"
	"testing"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/spenmo/entity"
	"github.com/indrasaputra/spenmo/internal/repository/postgres"
)

var (
	testCtx                = context.Background()
	errPostgresInternalMsg = "database down"
	errPostgresInternal    = errors.New(errPostgresInternalMsg)
)

type CardExecutor struct {
	card *postgres.Card
	pgx  pgxmock.PgxPoolIface
}

func TestNewCard(t *testing.T) {
	t.Run("successfully create an instance of Card", func(t *testing.T) {
		exec := createCardExecutor()
		assert.NotNil(t, exec.card)
	})
}

func TestCard_Insert(t *testing.T) {
	t.Run("nil card is prohibited", func(t *testing.T) {
		exec := createCardExecutor()

		err := exec.card.Insert(testCtx, nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyCard(), err)
	})

	t.Run("postgres database returns internal error", func(t *testing.T) {
		exec := createCardExecutor()
		exec.pgx.
			ExpectExec(`INSERT INTO user_cards \(user_id, wallet_id, limit_daily, limit_monthly, created_at, updated_at\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6\)`).
			WillReturnError(errPostgresInternal)

		err := exec.card.Insert(testCtx, createCard())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(errPostgresInternalMsg), err)
	})

	t.Run("success insert a new card", func(t *testing.T) {
		exec := createCardExecutor()
		exec.pgx.
			ExpectExec(`INSERT INTO user_cards \(user_id, wallet_id, limit_daily, limit_monthly, created_at, updated_at\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6\)`).
			WillReturnResult(pgxmock.NewResult("INSERT", 1))

		err := exec.card.Insert(testCtx, createCard())

		assert.Nil(t, err)
	})
}

func createCardExecutor() *CardExecutor {
	mock, err := pgxmock.NewPool(pgxmock.MonitorPingsOption(true))
	if err != nil {
		log.Panicf("error opening a stub database connection: %v\n", err)
	}

	card := postgres.NewCard(mock)
	return &CardExecutor{
		card: card,
		pgx:  mock,
	}
}

func createCard() *entity.UserCard {
	return &entity.UserCard{ID: 1, UserID: 1, WalletID: 1, LimitDaily: 1000000, LimitMonthly: 5000000}
}
