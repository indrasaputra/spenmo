package model_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"testing"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/spenmo/entity"
	"github.com/indrasaputra/spenmo/internal/repository/model"
	"github.com/indrasaputra/spenmo/internal/repository/model/ent"
)

var (
	testCtx                = context.Background()
	testUserID             = int64(1)
	testCardID             = int64(1)
	errPostgresInternalMsg = "database down"
	errPostgresInternal    = errors.New(errPostgresInternalMsg)
)

type CardExecutor struct {
	db     *sql.DB
	mock   sqlmock.Sqlmock
	client *ent.Client
}

func (c CardExecutor) close() {
	_ = c.db.Close()
	_ = c.client.Close()
}

func TestNewCard(t *testing.T) {
	t.Run("successfully create an instance of Card", func(t *testing.T) {
		card := model.NewCard(&ent.Client{})
		assert.NotNil(t, card)
	})
}

func TestCard_Insert(t *testing.T) {
	query := `INSERT INTO "user_cards" \("created_at", "updated_at", "limit_daily", "limit_monthly", "user_id", "wallet_id"\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6\) RETURNING "id"`

	t.Run("nil card is prohibited", func(t *testing.T) {
		exec := createCardExecutor()
		defer exec.close()
		card := model.NewCard(exec.client)

		err := card.Insert(testCtx, nil)

		assert.NotNil(t, err)
	})

	t.Run("database returns error", func(t *testing.T) {
		exec := createCardExecutor()
		defer exec.close()
		exec.mock.ExpectBegin()
		exec.mock.ExpectQuery(query).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(errPostgresInternal)
		exec.mock.ExpectRollback()
		card := model.NewCard(exec.client)

		err := card.Insert(testCtx, createCard())

		assert.NotNil(t, err)
	})

	t.Run("successfully insert card to database", func(t *testing.T) {
		exec := createCardExecutor()
		defer exec.close()
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

func TestCard_GetByID(t *testing.T) {
	query := `SELECT DISTINCT "user_cards"."id", "user_cards"."created_at", "user_cards"."updated_at", "user_cards"."deleted_at", "user_cards"."user_id", "user_cards"."wallet_id", "user_cards"."limit_daily", "user_cards"."limit_monthly" FROM "user_cards" WHERE \("user_cards"."id" = \$1 AND "user_cards"."user_id" = \$2\) AND "user_cards"."deleted_at" IS NULL LIMIT 2`

	t.Run("record not found", func(t *testing.T) {
		exec := createCardExecutor()
		defer exec.close()
		exec.mock.ExpectQuery(query).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(&ent.NotFoundError{})
		card := model.NewCard(exec.client)

		res, err := card.GetByID(testCtx, testUserID, testCardID)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound(), err)
		assert.Nil(t, res)
	})

	t.Run("database is down", func(t *testing.T) {
		exec := createCardExecutor()
		defer exec.close()
		exec.mock.ExpectQuery(query).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(errPostgresInternal)
		card := model.NewCard(exec.client)

		res, err := card.GetByID(testCtx, testUserID, testCardID)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(errPostgresInternalMsg), err)
		assert.Nil(t, res)
	})

	t.Run("successfully get record by id", func(t *testing.T) {
		exec := createCardExecutor()
		defer exec.close()
		exec.mock.ExpectQuery(query).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		card := model.NewCard(exec.client)

		res, err := card.GetByID(testCtx, testUserID, testCardID)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func TestCard_GetAll(t *testing.T) {
	query := `SELECT DISTINCT "user_cards"."id", "user_cards"."created_at", "user_cards"."updated_at", "user_cards"."deleted_at", "user_cards"."user_id", "user_cards"."wallet_id", "user_cards"."limit_daily", "user_cards"."limit_monthly" FROM "user_cards" WHERE "user_cards"."user_id" = \$1 AND "user_cards"."deleted_at" IS NULL`

	t.Run("database is down", func(t *testing.T) {
		exec := createCardExecutor()
		defer exec.close()
		exec.mock.ExpectQuery(query).
			WithArgs(sqlmock.AnyArg()).
			WillReturnError(errPostgresInternal)
		card := model.NewCard(exec.client)

		res, err := card.GetAll(testCtx, testUserID)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(errPostgresInternalMsg), err)
		assert.Nil(t, res)
	})

	t.Run("successfully get all records owned by certain user id", func(t *testing.T) {
		exec := createCardExecutor()
		defer exec.close()
		exec.mock.ExpectQuery(query).
			WithArgs(sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		card := model.NewCard(exec.client)

		res, err := card.GetAll(testCtx, testUserID)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.NotEmpty(t, res)
	})
}

func TestCard_Update(t *testing.T) {
	query := `UPDATE "user_cards" SET "updated_at" = \$1, "limit_daily" = \$2, "limit_monthly" = \$3 WHERE \("user_cards"."id" = \$4 AND "user_cards"."user_id" = \$5\) AND "user_cards"."deleted_at" IS NULL`

	t.Run("database is down", func(t *testing.T) {
		exec := createCardExecutor()
		defer exec.close()
		exec.mock.ExpectBegin()
		exec.mock.ExpectExec(query).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(errPostgresInternal)
		exec.mock.ExpectRollback()
		card := model.NewCard(exec.client)

		err := card.Update(testCtx, createCard())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(errPostgresInternalMsg), err)
	})

	t.Run("record not found", func(t *testing.T) {
		exec := createCardExecutor()
		defer exec.close()
		exec.mock.ExpectBegin()
		exec.mock.ExpectExec(query).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(0, 0))
		exec.mock.ExpectCommit()
		card := model.NewCard(exec.client)

		err := card.Update(testCtx, createCard())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound(), err)
	})

	t.Run("successfully update record", func(t *testing.T) {
		exec := createCardExecutor()
		defer exec.close()
		exec.mock.ExpectBegin()
		exec.mock.ExpectExec(query).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		exec.mock.ExpectCommit()
		card := model.NewCard(exec.client)

		err := card.Update(testCtx, createCard())

		assert.Nil(t, err)
	})
}

func TestCard_Delete(t *testing.T) {
	query := `DELETE FROM "user_cards" WHERE \("user_cards"."id" = \$1 AND "user_cards"."user_id" = \$2\) AND "user_cards"."deleted_at" IS NULL`

	t.Run("database is down", func(t *testing.T) {
		exec := createCardExecutor()
		defer exec.close()
		exec.mock.ExpectBegin()
		exec.mock.ExpectExec(query).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(errPostgresInternal)
		exec.mock.ExpectRollback()
		card := model.NewCard(exec.client)

		err := card.Delete(testCtx, testUserID, testCardID)

		fmt.Println(err)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(errPostgresInternalMsg), err)
	})

	t.Run("record not found", func(t *testing.T) {
		exec := createCardExecutor()
		defer exec.close()
		exec.mock.ExpectBegin()
		exec.mock.ExpectExec(query).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(0, 0))
		exec.mock.ExpectCommit()
		card := model.NewCard(exec.client)

		err := card.Delete(testCtx, testUserID, testCardID)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound(), err)
	})

	t.Run("successfully delete record", func(t *testing.T) {
		exec := createCardExecutor()
		defer exec.close()
		exec.mock.ExpectBegin()
		exec.mock.ExpectExec(query).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		exec.mock.ExpectCommit()
		card := model.NewCard(exec.client)

		err := card.Delete(testCtx, testUserID, testCardID)

		assert.Nil(t, err)
	})
}

func createCardExecutor() *CardExecutor {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("error initiate mock: %v", err)
	}
	driver := entsql.OpenDB(dialect.Postgres, db)
	entClient := ent.NewClient(ent.Driver(driver))

	return &CardExecutor{db, mock, entClient}
}

func createCard() *entity.UserCard {
	return &entity.UserCard{ID: 1, UserID: 1, WalletID: 1, LimitDaily: 1000000, LimitMonthly: 5000000}
}
