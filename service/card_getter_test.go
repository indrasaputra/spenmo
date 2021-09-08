package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/spenmo/entity"
	"github.com/indrasaputra/spenmo/service"
	mock_service "github.com/indrasaputra/spenmo/test/mock/service"
)

var (
	testCardID = int64(1)
	testUserID = int64(1)
)

type CardGetterExecutor struct {
	getter *service.CardGetter
	repo   *mock_service.MockGetCardRepository
}

func TestNewCardGetter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of CardGetter", func(t *testing.T) {
		exec := createCardGetterExecutor(ctrl)
		assert.NotNil(t, exec.getter)
	})
}

func TestCardGetter_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("repository returns internal error", func(t *testing.T) {
		exec := createCardGetterExecutor(ctrl)
		exec.repo.EXPECT().GetByID(testCtx, testUserID, testCardID).Return(nil, entity.ErrInternal(""))

		res, err := exec.getter.GetByID(testCtx, testUserID, testCardID)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(""), err)
		assert.Nil(t, res)
	})

	t.Run("repository returns not found error", func(t *testing.T) {
		exec := createCardGetterExecutor(ctrl)
		exec.repo.EXPECT().GetByID(testCtx, testUserID, testCardID).Return(nil, entity.ErrNotFound())

		res, err := exec.getter.GetByID(testCtx, testUserID, testCardID)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound(), err)
		assert.Nil(t, res)
	})

	t.Run("successfully get a single card", func(t *testing.T) {
		exec := createCardGetterExecutor(ctrl)
		card := createCard()
		exec.repo.EXPECT().GetByID(testCtx, testUserID, testCardID).Return(card, nil)

		res, err := exec.getter.GetByID(testCtx, testUserID, testCardID)

		assert.Nil(t, err)
		assert.Equal(t, card, res)
	})
}

func TestCardGetter_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("repository returns internal error", func(t *testing.T) {
		exec := createCardGetterExecutor(ctrl)
		exec.repo.EXPECT().GetAll(testCtx, testUserID).Return([]*entity.UserCard{}, entity.ErrInternal(""))

		res, err := exec.getter.GetAll(testCtx, testUserID)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(""), err)
		assert.Empty(t, res)
	})

	t.Run("repository returns empty list", func(t *testing.T) {
		exec := createCardGetterExecutor(ctrl)
		exec.repo.EXPECT().GetAll(testCtx, testUserID).Return([]*entity.UserCard{}, nil)

		res, err := exec.getter.GetAll(testCtx, testUserID)

		assert.Nil(t, err)
		assert.Empty(t, res)
	})

	t.Run("successfully get cards", func(t *testing.T) {
		exec := createCardGetterExecutor(ctrl)
		exec.repo.EXPECT().GetAll(testCtx, testUserID).Return([]*entity.UserCard{createCard()}, nil)

		res, err := exec.getter.GetAll(testCtx, testUserID)

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
	})
}

func createCardGetterExecutor(ctrl *gomock.Controller) *CardGetterExecutor {
	r := mock_service.NewMockGetCardRepository(ctrl)
	g := service.NewCardGetter(r)
	return &CardGetterExecutor{
		getter: g,
		repo:   r,
	}
}
