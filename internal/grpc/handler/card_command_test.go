package handler_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/indrasaputra/hashids"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/spenmo/entity"
	"github.com/indrasaputra/spenmo/internal/grpc/handler"
	"github.com/indrasaputra/spenmo/internal/grpc/interceptor"
	api "github.com/indrasaputra/spenmo/proto/indrasaputra/spenmo/v1"
	mock_service "github.com/indrasaputra/spenmo/test/mock/service"
)

var (
	testCtx = context.WithValue(context.Background(), interceptor.ContextKeyUser, int64(1))
)

type CardCommandExecutor struct {
	handler *handler.CardCommand

	creator *mock_service.MockCreateCard
	updater *mock_service.MockUpdateCard
	deleter *mock_service.MockDeleteCard
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

func TestCardCommand_UpdateCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("nil request is prohibited", func(t *testing.T) {
		exec := createCardCommandExecutor(ctrl)

		res, err := exec.handler.UpdateCard(testCtx, nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyCard(), err)
		assert.Nil(t, res)
	})

	t.Run("empty card is prohibited", func(t *testing.T) {
		exec := createCardCommandExecutor(ctrl)

		res, err := exec.handler.UpdateCard(testCtx, &api.UpdateCardRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyCard(), err)
		assert.Nil(t, res)
	})

	t.Run("card id is not hashid", func(t *testing.T) {
		exec := createCardCommandExecutor(ctrl)
		req := createUpdateCardRequest()
		req.Id = "abc"

		res, err := exec.handler.UpdateCard(testCtx, req)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInvalidID(), err)
		assert.Nil(t, res)
	})

	t.Run("card creator service returns error", func(t *testing.T) {
		exec := createCardCommandExecutor(ctrl)
		exec.updater.EXPECT().Update(testCtx, gomock.Any()).Return(entity.ErrInternal(""))

		res, err := exec.handler.UpdateCard(testCtx, createUpdateCardRequest())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(""), err)
		assert.Nil(t, res)
	})

	t.Run("success update user's card", func(t *testing.T) {
		exec := createCardCommandExecutor(ctrl)
		exec.updater.EXPECT().Update(testCtx, gomock.Any()).Return(nil)

		res, err := exec.handler.UpdateCard(testCtx, createUpdateCardRequest())

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func TestCardCommand_DeleteCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("nil request is prohibited", func(t *testing.T) {
		exec := createCardCommandExecutor(ctrl)

		res, err := exec.handler.DeleteCard(testCtx, nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyCard(), err)
		assert.Nil(t, res)
	})

	t.Run("card id is not hashid", func(t *testing.T) {
		exec := createCardCommandExecutor(ctrl)
		req := createDeleteCardRequest()
		req.Id = "abc"

		res, err := exec.handler.DeleteCard(testCtx, req)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInvalidID(), err)
		assert.Nil(t, res)
	})

	t.Run("card creator service returns error", func(t *testing.T) {
		exec := createCardCommandExecutor(ctrl)
		card := createCard()
		exec.deleter.EXPECT().Delete(testCtx, card.UserID, card.ID).Return(entity.ErrInternal(""))

		res, err := exec.handler.DeleteCard(testCtx, createDeleteCardRequest())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(""), err)
		assert.Nil(t, res)
	})

	t.Run("success delete user's card", func(t *testing.T) {
		exec := createCardCommandExecutor(ctrl)
		card := createCard()
		exec.deleter.EXPECT().Delete(testCtx, card.UserID, card.ID).Return(nil)

		res, err := exec.handler.DeleteCard(testCtx, createDeleteCardRequest())

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
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

func createUpdateCardRequest() *api.UpdateCardRequest {
	return &api.UpdateCardRequest{
		Card: createCardAPI(),
	}
}

func createDeleteCardRequest() *api.DeleteCardRequest {
	return &api.DeleteCardRequest{
		Id: createCardAPI().GetId(),
	}
}

func createCardCommandExecutor(ctrl *gomock.Controller) *CardCommandExecutor {
	c := mock_service.NewMockCreateCard(ctrl)
	u := mock_service.NewMockUpdateCard(ctrl)
	d := mock_service.NewMockDeleteCard(ctrl)
	h := handler.NewCardCommand(c, u, d)
	return &CardCommandExecutor{
		handler: h,
		creator: c,
		updater: u,
		deleter: d,
	}
}
