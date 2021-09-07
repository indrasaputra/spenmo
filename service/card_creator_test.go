package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/spenmo/entity"
	"github.com/indrasaputra/spenmo/service"
	mock_service "github.com/indrasaputra/spenmo/test/mock/service"
)

var (
	testCtx = context.Background()
)

type CardCreatorExecutor struct {
	creator *service.CardCreator
	repo    *mock_service.MockCreateCardRepository
}

func TestNewCardCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of CardCreator", func(t *testing.T) {
		exec := createCardCreatorExecutor(ctrl)
		assert.NotNil(t, exec.creator)
	})
}

func TestCardCreator_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("empty card is prohibited", func(t *testing.T) {
		exec := createCardCreatorExecutor(ctrl)

		err := exec.creator.Create(testCtx, nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyCard(), err)
	})

	t.Run("card without user is prohibited", func(t *testing.T) {
		exec := createCardCreatorExecutor(ctrl)

		err := exec.creator.Create(testCtx, createCardWithInvalidUser())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInvalidUser(), err)
	})

	t.Run("card without wallet is prohibited", func(t *testing.T) {
		exec := createCardCreatorExecutor(ctrl)

		err := exec.creator.Create(testCtx, createCardWithInvalidWallet())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInvalidWallet(), err)
	})

	t.Run("card with limit daily = 0 is prohibited", func(t *testing.T) {
		exec := createCardCreatorExecutor(ctrl)

		err := exec.creator.Create(testCtx, createCardWithInvalidLimitDaily())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInvalidLimit(), err)
	})

	t.Run("card with limit monthly = 0 is prohibited", func(t *testing.T) {
		exec := createCardCreatorExecutor(ctrl)

		err := exec.creator.Create(testCtx, createCardWithInvalidLimitMonthly())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInvalidLimit(), err)
	})

	t.Run("repo returns error", func(t *testing.T) {
		exec := createCardCreatorExecutor(ctrl)
		card := createCard()
		exec.repo.EXPECT().Insert(testCtx, card).Return(entity.ErrInternal(""))

		err := exec.creator.Create(testCtx, card)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(""), err)
	})

	t.Run("success create a card", func(t *testing.T) {
		exec := createCardCreatorExecutor(ctrl)
		card := createCard()
		exec.repo.EXPECT().Insert(testCtx, card).Return(nil)

		err := exec.creator.Create(testCtx, card)

		assert.Nil(t, err)
	})
}

func createCard() *entity.UserCard {
	return &entity.UserCard{ID: 1, UserID: 1, WalletID: 1, LimitDaily: 1000000, LimitMonthly: 5000000}
}

func createCardWithInvalidUser() *entity.UserCard {
	card := createCard()
	card.UserID = 0
	return card
}

func createCardWithInvalidWallet() *entity.UserCard {
	card := createCard()
	card.WalletID = 0
	return card
}

func createCardWithInvalidLimitDaily() *entity.UserCard {
	card := createCard()
	card.LimitDaily = 0
	return card
}

func createCardWithInvalidLimitMonthly() *entity.UserCard {
	card := createCard()
	card.LimitMonthly = 0
	return card
}

func createCardCreatorExecutor(ctrl *gomock.Controller) *CardCreatorExecutor {
	r := mock_service.NewMockCreateCardRepository(ctrl)
	c := service.NewCardCreator(r)
	return &CardCreatorExecutor{
		creator: c,
		repo:    r,
	}
}
