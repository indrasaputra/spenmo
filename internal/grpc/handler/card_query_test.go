package handler_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/spenmo/entity"
	"github.com/indrasaputra/spenmo/internal/grpc/handler"
	api "github.com/indrasaputra/spenmo/proto/indrasaputra/spenmo/v1"
	mock_service "github.com/indrasaputra/spenmo/test/mock/service"
)

type CardQueryExecutor struct {
	handler *handler.CardQuery

	getter *mock_service.MockGetCard
}

func TestNewCardQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successful create an instance of CardQuery", func(t *testing.T) {
		exec := createCardQueryExecutor(ctrl)
		assert.NotNil(t, exec.handler)
	})
}

func TestCardQuery_GetCardByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("nil request is prohibited", func(t *testing.T) {
		exec := createCardQueryExecutor(ctrl)

		res, err := exec.handler.GetCardByID(testCtx, nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyCard(), err)
		assert.Nil(t, res)
	})

	t.Run("card is not hashid", func(t *testing.T) {
		exec := createCardQueryExecutor(ctrl)

		req := createGetCardByIDRequest()
		req.Id = "abc"
		res, err := exec.handler.GetCardByID(testCtx, req)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInvalidID(), err)
		assert.Nil(t, res)
	})

	t.Run("card not found", func(t *testing.T) {
		exec := createCardQueryExecutor(ctrl)
		card := createCard()
		exec.getter.EXPECT().GetByID(testCtx, card.UserID, card.ID).Return(nil, entity.ErrNotFound())

		res, err := exec.handler.GetCardByID(testCtx, createGetCardByIDRequest())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound(), err)
		assert.Nil(t, res)
	})

	t.Run("getter service returns error", func(t *testing.T) {
		exec := createCardQueryExecutor(ctrl)
		card := createCard()
		exec.getter.EXPECT().GetByID(testCtx, card.UserID, card.ID).Return(nil, entity.ErrInternal(""))

		res, err := exec.handler.GetCardByID(testCtx, createGetCardByIDRequest())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(""), err)
		assert.Nil(t, res)
	})

	t.Run("success get a card", func(t *testing.T) {
		exec := createCardQueryExecutor(ctrl)
		card := createCard()
		exec.getter.EXPECT().GetByID(testCtx, card.UserID, card.ID).Return(createCard(), nil)

		res, err := exec.handler.GetCardByID(testCtx, createGetCardByIDRequest())

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func TestCardQuery_GetAllCards(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("nil request is prohibited", func(t *testing.T) {
		exec := createCardQueryExecutor(ctrl)

		res, err := exec.handler.GetAllCards(testCtx, nil)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrEmptyCard(), err)
		assert.Empty(t, res)
	})

	t.Run("getter service returns error", func(t *testing.T) {
		exec := createCardQueryExecutor(ctrl)
		card := createCard()
		exec.getter.EXPECT().GetAll(testCtx, card.UserID).Return([]*entity.UserCard{}, entity.ErrInternal(""))

		res, err := exec.handler.GetAllCards(testCtx, createGetAllCardsRequest())

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrInternal(""), err)
		assert.Empty(t, res)
	})

	t.Run("getter service returns error", func(t *testing.T) {
		exec := createCardQueryExecutor(ctrl)
		card := createCard()
		exec.getter.EXPECT().GetAll(testCtx, card.UserID).Return([]*entity.UserCard{createCard()}, nil)

		res, err := exec.handler.GetAllCards(testCtx, createGetAllCardsRequest())

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
	})
}

func createCardQueryExecutor(ctrl *gomock.Controller) *CardQueryExecutor {
	g := mock_service.NewMockGetCard(ctrl)

	h := handler.NewCardQuery(g)
	return &CardQueryExecutor{
		handler: h,
		getter:  g,
	}
}

func createGetCardByIDRequest() *api.GetCardByIDRequest {
	return &api.GetCardByIDRequest{
		Id: createCardAPI().GetId(),
	}
}

func createGetAllCardsRequest() *api.GetAllCardsRequest {
	return &api.GetAllCardsRequest{}
}
