package service

import (
	"context"
)

// DeleteCard defines the interface to delete user's card.
type DeleteCard interface {
	// Delete deletes an existing user's card.
	Delete(ctx context.Context, userID, cardID int64) error
}

// DeleteCardRepository defines interface to delete the card in the storage.
type DeleteCardRepository interface {
	// Delete deletes a card in the storage.
	Delete(ctx context.Context, userID, cardID int64) error
}

// CardDeleter is responsible to delete a user's card.
type CardDeleter struct {
	repo DeleteCardRepository
}

// NewCardDeleter creates an instance of CardDeleter.
func NewCardDeleter(repo DeleteCardRepository) *CardDeleter {
	return &CardDeleter{repo: repo}
}

// Delete deletes an existing user's card.
func (cc *CardDeleter) Delete(ctx context.Context, userID, cardID int64) error {
	return cc.repo.Delete(ctx, userID, cardID)
}
