package postgres

import (
	"context"
	"time"

	"github.com/indrasaputra/spenmo/entity"
)

// Card is responsible to connect card entity with user_cards table in PostgreSQL.
type Card struct {
	pool PgxPoolIface
}

// NewCard creates an instance of Card.
func NewCard(pool PgxPoolIface) *Card {
	return &Card{pool: pool}
}

// Insert inserts the card into the user_cards table.
func (c *Card) Insert(ctx context.Context, card *entity.UserCard) error {
	if card == nil {
		return entity.ErrEmptyCard()
	}
	card.CreatedAt = time.Now().UTC()
	card.UpdatedAt = time.Now().UTC()
	card.DeletedAt = nil

	query := "INSERT INTO " +
		"user_cards (user_id, wallet_id, limit_daily, limit_monthly, created_at, updated_at) " +
		"VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := c.pool.Exec(ctx, query,
		card.UserID,
		card.WalletID,
		card.LimitDaily,
		card.LimitMonthly,
		card.CreatedAt,
		card.UpdatedAt,
	)

	if err != nil {
		return entity.ErrInternal(err.Error())
	}
	return nil
}
