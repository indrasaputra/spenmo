package handler_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/indrasaputra/hashids"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/spenmo/entity"
	"github.com/indrasaputra/spenmo/internal/grpc/handler"
	api "github.com/indrasaputra/spenmo/proto/indrasaputra/spenmo/v1"
	mock_service "github.com/indrasaputra/spenmo/test/mock/service"
)

var (
	testCtx = context.Background()
)

type CardCommandExecutor struct {
	handler *handler.CardCommand
	creator *mock_service.MockCreateCard
}

func TestNewCardCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successful create an instance of CardCommand", func(t *testing.T) {
		exec := createCardCommandExecutor(ctrl)
		assert.NotNil(t, exec.handler)
	})
}

func TestCardCommand_CreateCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("nil request is prohibited", func(t *testing.T) {
		exec := createCardCommandExecutor(ctrl)

		res, err := exec.handler.CreateCard(testCtx, nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyCard(), err)
		assert.Nil(t, res)
	})

	t.Run("empty card is prohibited", func(t *testing.T) {
		exec := createCardCommandExecutor(ctrl)

		res, err := exec.handler.CreateCard(testCtx, &api.CreateCardRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyCard(), err)
		assert.Nil(t, res)
	})

	t.Run("card with invalid wallet is prohibited", func(t *testing.T) {
		exec := createCardCommandExecutor(ctrl)
		request := createCreateCardRequest()
		request.Card.WalletId = "abc"

		res, err := exec.handler.CreateCard(testCtx, request)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInvalidWallet(), err)
		assert.Nil(t, res)
	})

	t.Run("card creator service returns error", func(t *testing.T) {
		exec := createCardCommandExecutor(ctrl)
		exec.creator.EXPECT().Create(testCtx, gomock.Any()).Return(entity.ErrInternal(""))

		res, err := exec.handler.CreateCard(testCtx, createCreateCardRequest())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(""), err)
		assert.Nil(t, res)
	})

	t.Run("success create user's card", func(t *testing.T) {
		exec := createCardCommandExecutor(ctrl)
		exec.creator.EXPECT().Create(testCtx, gomock.Any()).Return(nil)

		res, err := exec.handler.CreateCard(testCtx, createCreateCardRequest())

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func createCardCommandExecutor(ctrl *gomock.Controller) *CardCommandExecutor {
	c := mock_service.NewMockCreateCard(ctrl)
	h := handler.NewCardCommand(c)
	return &CardCommandExecutor{
		handler: h,
		creator: c,
	}
}

func createCard() *entity.UserCard {
	return &entity.UserCard{ID: 1, UserID: 1, WalletID: 1, LimitDaily: 1000000, LimitMonthly: 5000000}
}

func createCardAPI() *api.Card {
	card := createCard()
	id := hashids.ID(card.ID).EncodeString()
	userID := hashids.ID(card.UserID).EncodeString()
	walletID := hashids.ID(card.WalletID).EncodeString()

	return &api.Card{
		Id:           id,
		UserId:       userID,
		WalletId:     walletID,
		LimitDaily:   card.LimitDaily,
		LimitMonthly: card.LimitMonthly,
	}
}

func createCreateCardRequest() *api.CreateCardRequest {
	return &api.CreateCardRequest{
		Card: createCardAPI(),
	}
}
