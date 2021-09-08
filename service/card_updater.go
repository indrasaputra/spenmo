package service

import (
	"context"

	"github.com/indrasaputra/spenmo/entity"
)

// UpdateCard defines the interface to update user's card.
type UpdateCard interface {
	// Update updates an existing user's card.
	Update(ctx context.Context, card *entity.UserCard) error
}

// UpdateCardRepository defines interface to update the card in the storage.
type UpdateCardRepository interface {
	// Update updates a card in the storage.
	Update(ctx context.Context, card *entity.UserCard) error
}

// CardUpdater is responsible to update a user's card.
type CardUpdater struct {
	repo UpdateCardRepository
}

// NewCardUpdater creates an instance of CardUpdater.
func NewCardUpdater(repo UpdateCardRepository) *CardUpdater {
	return &CardUpdater{repo: repo}
}

// Update updates an existing user's card.
func (cu *CardUpdater) Update(ctx context.Context, card *entity.UserCard) error {
	if err := validateCardForUpdate(card); err != nil {
		return err
	}
	return cu.repo.Update(ctx, card)
}

func validateCardForUpdate(card *entity.UserCard) error {
	if card == nil {
		return entity.ErrEmptyCard()
	}
	if card.UserID == 0 {
		return entity.ErrInvalidUser()
	}
	if card.LimitDaily <= 0 || card.LimitMonthly <= 0 {
		return entity.ErrInvalidLimit()
	}
	return nil
}
