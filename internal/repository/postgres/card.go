package postgres

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4"

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

// GetByID gets a single user's card from the repository.
// If the user's card can't be found, it returns NotFound error.
func (c *Card) GetByID(ctx context.Context, userID, cardID int64) (*entity.UserCard, error) {
	query := "SELECT id, user_id, wallet_id, limit_daily, limit_monthly, created_at, updated_at FROM user_cards WHERE id = $1 AND user_id = $2 AND deleted_at IS NULL LIMIT 1"
	row := c.pool.QueryRow(ctx, query, cardID, userID)

	res := entity.UserCard{}
	err := row.Scan(&res.ID, &res.UserID, &res.WalletID, &res.LimitDaily, &res.LimitMonthly, &res.CreatedAt, &res.UpdatedAt)
	if err == pgx.ErrNoRows {
		return nil, entity.ErrNotFound()
	}
	if err != nil {
		return nil, entity.ErrInternal(err.Error())
	}
	return &res, nil
}

// GetAll gets all user's cards available in repository.
// If there isn't any user's card in repository, it returns empty list of user's card and nil error.
func (c *Card) GetAll(ctx context.Context, userID int64) ([]*entity.UserCard, error) {
	query := "SELECT id, user_id, wallet_id, limit_daily, limit_monthly, created_at, updated_at FROM user_cards WHERE user_id = $1 AND deleted_at IS NULL"
	rows, err := c.pool.Query(ctx, query, userID)
	if err != nil {
		return []*entity.UserCard{}, entity.ErrInternal(err.Error())
	}
	defer rows.Close()

	res := []*entity.UserCard{}
	for rows.Next() {
		var tmp entity.UserCard
		if err := rows.Scan(&tmp.ID, &tmp.UserID, &tmp.WalletID, &tmp.LimitDaily, &tmp.LimitMonthly, &tmp.CreatedAt, &tmp.UpdatedAt); err != nil {
			log.Printf("[Card-GetAll] scan rows error: %s", err.Error())
			continue
		}
		res = append(res, &tmp)
	}
	if rows.Err() != nil {
		return []*entity.UserCard{}, entity.ErrInternal(rows.Err().Error())
	}
	return res, nil
}

// Update updates the given card information in repository.
func (c *Card) Update(ctx context.Context, card *entity.UserCard) error {
	query := "UPDATE user_cards " +
		"SET limit_daily = $1, limit_monthly = $2, updated_at = $3 " +
		"WHERE id = $4 AND user_id = $5 AND deleted_at IS NULL"

	tag, err := c.pool.Exec(ctx, query, card.LimitDaily, card.LimitMonthly, time.Now().UTC(), card.ID, card.UserID)
	if err != nil {
		return entity.ErrInternal(err.Error())
	}
	if tag.RowsAffected() == 0 {
		return entity.ErrNotFound()
	}
	return nil
}

// Delete deletes the given card information in repository.
// It performs soft delete.
func (c *Card) Delete(ctx context.Context, userID, cardID int64) error {
	query := "UPDATE user_cards " +
		"SET deleted_at = $1 " +
		"WHERE id = $2 AND user_id = $3 AND deleted_at IS NULL"

	tag, err := c.pool.Exec(ctx, query, time.Now().UTC(), cardID, userID)
	if err != nil {
		return entity.ErrInternal(err.Error())
	}
	if tag.RowsAffected() == 0 {
		return entity.ErrNotFound()
	}
	return nil
}
