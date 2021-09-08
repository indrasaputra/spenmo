package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/spenmo/entity"
	"github.com/indrasaputra/spenmo/service"
	mock_service "github.com/indrasaputra/spenmo/test/mock/service"
)

type CardUpdaterExecutor struct {
	updater *service.CardUpdater
	repo    *mock_service.MockUpdateCardRepository
}

func TestNewCardUpdater(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of CardUpdater", func(t *testing.T) {
		exec := createCardUpdaterExecutor(ctrl)
		assert.NotNil(t, exec.updater)
	})
}

func TestCardUpdater_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("empty card is prohibited", func(t *testing.T) {
		exec := createCardUpdaterExecutor(ctrl)

		err := exec.updater.Update(testCtx, nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyCard(), err)
	})

	t.Run("card without user is prohibited", func(t *testing.T) {
		exec := createCardUpdaterExecutor(ctrl)

		err := exec.updater.Update(testCtx, createCardWithInvalidUser())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInvalidUser(), err)
	})

	t.Run("card with limit daily = 0 is prohibited", func(t *testing.T) {
		exec := createCardUpdaterExecutor(ctrl)

		err := exec.updater.Update(testCtx, createCardWithInvalidLimitDaily())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInvalidLimit(), err)
	})

	t.Run("card with limit monthly = 0 is prohibited", func(t *testing.T) {
		exec := createCardUpdaterExecutor(ctrl)

		err := exec.updater.Update(testCtx, createCardWithInvalidLimitMonthly())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInvalidLimit(), err)
	})

	t.Run("card not found", func(t *testing.T) {
		exec := createCardUpdaterExecutor(ctrl)
		card := createCard()
		exec.repo.EXPECT().Update(testCtx, card).Return(entity.ErrNotFound())

		err := exec.updater.Update(testCtx, card)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound(), err)
	})

	t.Run("repo returns error", func(t *testing.T) {
		exec := createCardUpdaterExecutor(ctrl)
		card := createCard()
		exec.repo.EXPECT().Update(testCtx, card).Return(entity.ErrInternal(""))

		err := exec.updater.Update(testCtx, card)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(""), err)
	})

	t.Run("success update a card", func(t *testing.T) {
		exec := createCardUpdaterExecutor(ctrl)
		card := createCard()
		exec.repo.EXPECT().Update(testCtx, card).Return(nil)

		err := exec.updater.Update(testCtx, card)

		assert.Nil(t, err)
	})
}

func createCardUpdaterExecutor(ctrl *gomock.Controller) *CardUpdaterExecutor {
	r := mock_service.NewMockUpdateCardRepository(ctrl)
	u := service.NewCardUpdater(r)
	return &CardUpdaterExecutor{
		updater: u,
		repo:    r,
	}
}
