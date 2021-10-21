package model

import (
	"context"

	"github.com/indrasaputra/spenmo/entity"
	"github.com/indrasaputra/spenmo/internal/repository/model/ent"
	"github.com/indrasaputra/spenmo/internal/repository/model/ent/usercard"
)

// Card is responsible to connect card entity with user_cards table in PostgreSQL.
type Card struct {
	client *ent.Client
}

// NewCard creates an instance of Card.
func NewCard(client *ent.Client) *Card {
	return &Card{client: client}
}

// Insert inserts the card into the user_cards table.
func (c *Card) Insert(ctx context.Context, card *entity.UserCard) error {
	if card == nil {
		return entity.ErrEmptyCard()
	}

	_, err := c.client.UserCard.Create().
		SetUserID(card.UserID).
		SetWalletID(card.WalletID).
		SetLimitDaily(card.LimitDaily).
		SetLimitMonthly(card.LimitMonthly).
		Save(ctx)

	if err != nil {
		return entity.ErrInternal(err.Error())
	}
	return nil
}

// GetByID gets a single user's card from the repository.
// If the user's card can't be found, it returns NotFound error.
func (c *Card) GetByID(ctx context.Context, userID, cardID int64) (*entity.UserCard, error) {
	res, err := c.client.UserCard.
		Query().
		Where(
			usercard.ID(cardID),
			usercard.UserID(userID),
			usercard.DeletedAtIsNil(),
		).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, entity.ErrNotFound()
	}
	if err != nil {
		return nil, entity.ErrInternal(err.Error())
	}
	return userCardModelToEntity(res), nil
}

// GetAll gets all user's cards available in repository.
// If there isn't any user's card in repository, it returns empty list of user's card and nil error.
func (c *Card) GetAll(ctx context.Context, userID int64) ([]*entity.UserCard, error) {
	res, err := c.client.UserCard.
		Query().
		Where(
			usercard.UserID(userID),
			usercard.DeletedAtIsNil(),
		).
		All(ctx)

	if err != nil {
		return nil, entity.ErrInternal(err.Error())
	}
	return userCardListModelToEntity(res), nil
}

// Update updates the given card information in repository.
func (c *Card) Update(ctx context.Context, card *entity.UserCard) error {
	affected, err := c.client.UserCard.
		Update().
		Where(
			usercard.ID(card.ID),
			usercard.UserID(card.UserID),
			usercard.DeletedAtIsNil(),
		).
		SetLimitDaily(card.LimitDaily).
		SetLimitMonthly(card.LimitMonthly).
		Save(ctx)

	if err != nil {
		return entity.ErrInternal(err.Error())
	}
	if affected == 0 {
		return entity.ErrNotFound()
	}
	return nil
}

// Delete deletes the given card information in repository.
// It performs soft delete.
func (c *Card) Delete(ctx context.Context, userID, cardID int64) error {
	affected, err := c.client.UserCard.
		Delete().
		Where(
			usercard.ID(cardID),
			usercard.UserID(userID),
			usercard.DeletedAtIsNil(),
		).Exec(ctx)
	if err != nil {
		return entity.ErrInternal(err.Error())
	}
	if affected == 0 {
		return entity.ErrNotFound()
	}
	return nil
}

func userCardListModelToEntity(cards []*ent.UserCard) []*entity.UserCard {
	res := []*entity.UserCard{}
	for _, card := range cards {
		res = append(res, userCardModelToEntity(card))
	}
	return res
}

func userCardModelToEntity(card *ent.UserCard) *entity.UserCard {
	return &entity.UserCard{
		ID:           card.ID,
		UserID:       card.UserID,
		WalletID:     card.WalletID,
		LimitDaily:   card.LimitDaily,
		LimitMonthly: card.LimitMonthly,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
	}
}
