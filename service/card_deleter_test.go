package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/spenmo/entity"
	"github.com/indrasaputra/spenmo/service"
	mock_service "github.com/indrasaputra/spenmo/test/mock/service"
)

type CardDeleterExecutor struct {
	deleter *service.CardDeleter
	repo    *mock_service.MockDeleteCardRepository
}

func TestNewCardDeleter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of CardDeleter", func(t *testing.T) {
		exec := createCardDeleterExecutor(ctrl)
		assert.NotNil(t, exec.deleter)
	})
}

func TestCardDeleter_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("card not found", func(t *testing.T) {
		exec := createCardDeleterExecutor(ctrl)
		card := createCard()
		exec.repo.EXPECT().Delete(testCtx, card.UserID, card.ID).Return(entity.ErrNotFound())

		err := exec.deleter.Delete(testCtx, card.UserID, card.ID)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound(), err)
	})

	t.Run("repo returns error", func(t *testing.T) {
		exec := createCardDeleterExecutor(ctrl)
		card := createCard()
		exec.repo.EXPECT().Delete(testCtx, card.UserID, card.ID).Return(entity.ErrInternal(""))

		err := exec.deleter.Delete(testCtx, card.UserID, card.ID)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(""), err)
	})

	t.Run("success delete a card", func(t *testing.T) {
		exec := createCardDeleterExecutor(ctrl)
		card := createCard()
		exec.repo.EXPECT().Delete(testCtx, card.UserID, card.ID).Return(nil)

		err := exec.deleter.Delete(testCtx, card.UserID, card.ID)

		assert.Nil(t, err)
	})
}

func createCardDeleterExecutor(ctrl *gomock.Controller) *CardDeleterExecutor {
	r := mock_service.NewMockDeleteCardRepository(ctrl)
	d := service.NewCardDeleter(r)
	return &CardDeleterExecutor{
		deleter: d,
		repo:    r,
	}
}
