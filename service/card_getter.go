package service

import (
	"context"

	"github.com/indrasaputra/spenmo/entity"
)

// GetCard defines the interface to get user's card.
type GetCard interface {
	// GetByID gets a single user's card by its id.
	GetByID(ctx context.Context, userID, cardID int64) (*entity.UserCard, error)
	// GetAll gets all user's cards available in system.
	GetAll(ctx context.Context, userID int64) ([]*entity.UserCard, error)
}

// GetCardRepository defines the interface to get user's card from the repository.
type GetCardRepository interface {
	// GetByID gets a single user's card from the repository.
	// If the user's card can't be found, it returns NotFound error.
	GetByID(ctx context.Context, userID, cardID int64) (*entity.UserCard, error)
	// GetAll gets all user's cards available in repository.
	// If there isn't any user's card in repository, it returns empty list of user's card and nil error.
	GetAll(ctx context.Context, userID int64) ([]*entity.UserCard, error)
}

// CardGetter is responsible for getting user's card.
type CardGetter struct {
	repo GetCardRepository
}

// NewCardGetter creates an instance of CardGetter.
func NewCardGetter(repo GetCardRepository) *CardGetter {
	return &CardGetter{repo: repo}
}

// GetByID gets a single user's card.
// If the user's card can't be found, it returns NotFound error.
func (cg *CardGetter) GetByID(ctx context.Context, userID, cardID int64) (*entity.UserCard, error) {
	return cg.repo.GetByID(ctx, userID, cardID)
}

// GetAll gets all user's cards available in repository.
// If there isn't any user's card in repository, it returns empty list of user's card and nil error.
func (cg *CardGetter) GetAll(ctx context.Context, userID int64) ([]*entity.UserCard, error) {
	return cg.repo.GetAll(ctx, userID)
}
