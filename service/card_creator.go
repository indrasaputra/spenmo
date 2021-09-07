package service

import (
	"context"

	"github.com/indrasaputra/spenmo/entity"
)

// CreateCard defines the interface to create user's card.
type CreateCard interface {
	// Create creates a new user's card and store it in the storage.
	Create(ctx context.Context, card *entity.UserCard) error
}

// CreateCardRepository defines interface to store the card in the storage.
type CreateCardRepository interface {
	// Insert inserts a card in the storage.
	Insert(ctx context.Context, card *entity.UserCard) error
}

// CardCreator is responsible to create a user's card.
type CardCreator struct {
	repo CreateCardRepository
}

// NewCardCreator creates an instance of CardCreator.
func NewCardCreator(repo CreateCardRepository) *CardCreator {
	return &CardCreator{
		repo: repo,
	}
}

// Create creates a new user's card and store it in the storage.
func (cc *CardCreator) Create(ctx context.Context, card *entity.UserCard) error {
	if err := validateCard(card); err != nil {
		return err
	}
	return cc.repo.Insert(ctx, card)
}

func validateCard(card *entity.UserCard) error {
	if card == nil {
		return entity.ErrEmptyCard()
	}
	if card.UserID == 0 {
		return entity.ErrInvalidUser()
	}
	if card.WalletID == 0 {
		return entity.ErrInvalidWallet()
	}
	if card.LimitDaily <= 0 || card.LimitMonthly <= 0 {
		return entity.ErrInvalidLimit()
	}
	return nil
}
