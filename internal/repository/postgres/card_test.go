package postgres_test

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/spenmo/entity"
	"github.com/indrasaputra/spenmo/internal/repository/postgres"
)

var (
	testCtx                = context.Background()
	testUserID             = int64(1)
	testCardID             = int64(1)
	testWalletID           = int64(1)
	testLimitDaily         = float64(1000000)
	testLimitMonthly       = float64(5000000)
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

func TestCard_GetByID(t *testing.T) {
	t.Run("record not found", func(t *testing.T) {
		exec := createCardExecutor()
		exec.pgx.
			ExpectQuery(`SELECT id, user_id, wallet_id, limit_daily, limit_monthly, created_at, updated_at FROM user_cards WHERE id = \$1 AND user_id = \$2 AND deleted_at IS NULL LIMIT 1`).
			WillReturnError(pgx.ErrNoRows)

		res, err := exec.card.GetByID(testCtx, testUserID, testCardID)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound(), err)
		assert.Nil(t, res)
	})

	t.Run("query returns error", func(t *testing.T) {
		exec := createCardExecutor()
		exec.pgx.
			ExpectQuery(`SELECT id, user_id, wallet_id, limit_daily, limit_monthly, created_at, updated_at FROM user_cards WHERE id = \$1 AND user_id = \$2 AND deleted_at IS NULL LIMIT 1`).
			WillReturnError(errPostgresInternal)

		res, err := exec.card.GetByID(testCtx, testUserID, testCardID)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(errPostgresInternalMsg), err)
		assert.Nil(t, res)
	})

	t.Run("successfully retrieve a record", func(t *testing.T) {
		exec := createCardExecutor()
		exec.pgx.
			ExpectQuery(`SELECT id, user_id, wallet_id, limit_daily, limit_monthly, created_at, updated_at FROM user_cards WHERE id = \$1 AND user_id = \$2 AND deleted_at IS NULL LIMIT 1`).
			WillReturnRows(pgxmock.
				NewRows([]string{"id", "user_id", "wallet_id", "limit_daily", "limit_monthly", "created_at", "updated_at"}).
				AddRow(testCardID, testUserID, testWalletID, testLimitDaily, testLimitMonthly, time.Now(), time.Now()),
			)

		res, err := exec.card.GetByID(testCtx, testUserID, testCardID)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func TestCard_GetAll(t *testing.T) {
	t.Run("query returns error", func(t *testing.T) {
		exec := createCardExecutor()
		exec.pgx.
			ExpectQuery(`SELECT id, user_id, wallet_id, limit_daily, limit_monthly, created_at, updated_at FROM user_cards WHERE user_id = \$1 AND deleted_at IS NULL`).
			WillReturnError(errPostgresInternal)

		res, err := exec.card.GetAll(testCtx, testUserID)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(errPostgresInternalMsg), err)
		assert.Empty(t, res)
	})

	t.Run("scan returns error", func(t *testing.T) {
		exec := createCardExecutor()
		exec.pgx.
			ExpectQuery(`SELECT id, user_id, wallet_id, limit_daily, limit_monthly, created_at, updated_at FROM user_cards WHERE user_id = \$1 AND deleted_at IS NULL`).
			WillReturnRows(pgxmock.
				NewRows([]string{"id", "user_id", "wallet_id", "limit_daily", "limit_monthly", "created_at", "updated_at"}).
				AddRow(testCardID, testUserID, testWalletID, testLimitDaily, testLimitMonthly, time.Now(), time.Now()).
				AddRow("#$1", testUserID, testWalletID, testLimitDaily, testLimitMonthly, "time.Now()", "time.Now()"),
			)

		res, err := exec.card.GetAll(testCtx, testUserID)

		assert.Nil(t, err)
		assert.Equal(t, 1, len(res))
	})

	t.Run("rows error after all scans", func(t *testing.T) {
		exec := createCardExecutor()
		exec.pgx.
			ExpectQuery(`SELECT id, user_id, wallet_id, limit_daily, limit_monthly, created_at, updated_at FROM user_cards WHERE user_id = \$1 AND deleted_at IS NULL`).
			WillReturnRows(pgxmock.
				NewRows([]string{"id", "user_id", "wallet_id", "limit_daily", "limit_monthly", "created_at", "updated_at"}).
				AddRow(testCardID, testUserID, testWalletID, testLimitDaily, testLimitMonthly, time.Now(), time.Now()).
				AddRow(testCardID, testUserID, testWalletID, testLimitDaily, testLimitMonthly, time.Now(), time.Now()).
				RowError(2, errPostgresInternal),
			)

		res, err := exec.card.GetAll(testCtx, testUserID)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(errPostgresInternalMsg), err)
		assert.Empty(t, res)
	})

	t.Run("success get all records", func(t *testing.T) {
		exec := createCardExecutor()
		exec.pgx.
			ExpectQuery(`SELECT id, user_id, wallet_id, limit_daily, limit_monthly, created_at, updated_at FROM user_cards WHERE user_id = \$1 AND deleted_at IS NULL`).
			WillReturnRows(pgxmock.
				NewRows([]string{"id", "user_id", "wallet_id", "limit_daily", "limit_monthly", "created_at", "updated_at"}).
				AddRow(testCardID, testUserID, testWalletID, testLimitDaily, testLimitMonthly, time.Now(), time.Now()).
				AddRow(testCardID, testUserID, testWalletID, testLimitDaily, testLimitMonthly, time.Now(), time.Now()),
			)

		res, err := exec.card.GetAll(testCtx, testUserID)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(res))
	})
}

func TestCard_Update(t *testing.T) {
	t.Run("query returns error", func(t *testing.T) {
		exec := createCardExecutor()
		exec.pgx.
			ExpectExec(`UPDATE user_cards SET limit_daily = \$1, limit_monthly = \$2, updated_at = \$3 WHERE id = \$4 AND user_id = \$5 AND deleted_at IS NULL`).
			WillReturnError(errPostgresInternal)

		err := exec.card.Update(testCtx, createCard())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(errPostgresInternalMsg), err)
	})

	t.Run("no card updated", func(t *testing.T) {
		exec := createCardExecutor()
		exec.pgx.
			ExpectExec(`UPDATE user_cards SET limit_daily = \$1, limit_monthly = \$2, updated_at = \$3 WHERE id = \$4 AND user_id = \$5 AND deleted_at IS NULL`).
			WillReturnResult(pgxmock.NewResult("UPDATE", 0))

		err := exec.card.Update(testCtx, createCard())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound(), err)
	})

	t.Run("success update card", func(t *testing.T) {
		exec := createCardExecutor()
		exec.pgx.
			ExpectExec(`UPDATE user_cards SET limit_daily = \$1, limit_monthly = \$2, updated_at = \$3 WHERE id = \$4 AND user_id = \$5 AND deleted_at IS NULL`).
			WillReturnResult(pgxmock.NewResult("UPDATE", 1))

		err := exec.card.Update(testCtx, createCard())

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
