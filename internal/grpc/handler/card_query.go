package handler

import (
	"context"

	"github.com/indrasaputra/hashids"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/indrasaputra/spenmo/entity"
	"github.com/indrasaputra/spenmo/internal/grpc/interceptor"
	api "github.com/indrasaputra/spenmo/proto/indrasaputra/spenmo/v1"
	"github.com/indrasaputra/spenmo/service"
)

// CardQuery handles HTTP/2 gRPC request for retrieve card .
type CardQuery struct {
	api.UnimplementedCardQueryServiceServer

	getter service.GetCard
}

// NewCardQuery creates an instance of CardQuery.
func NewCardQuery(getter service.GetCard) *CardQuery {
	return &CardQuery{getter: getter}
}

// GetCardByID handles HTTP/2 gRPC request similar to GET in HTTP/1.1.
// It gets a single user's card by its id.
func (cq *CardQuery) GetCardByID(ctx context.Context, request *api.GetCardByIDRequest) (*api.GetCardByIDResponse, error) {
	if request == nil {
		return nil, entity.ErrEmptyCard()
	}

	userID := ctx.Value(interceptor.ContextKeyUser).(int64)
	cardID, err := hashids.DecodeHash([]byte(request.GetId()))
	if err != nil {
		return nil, entity.ErrInvalidID()
	}
	card, err := cq.getter.GetByID(ctx, userID, int64(cardID))
	if err != nil {
		return nil, err
	}
	return createGetCardByIDResponse(card), nil
}

// GetAllCards handles HTTP/2 gRPC request similar to GET in HTTP/1.1.
func (cq *CardQuery) GetAllCards(ctx context.Context, request *api.GetAllCardsRequest) (*api.GetAllCardsResponse, error) {
	if request == nil {
		return nil, entity.ErrEmptyCard()
	}

	userID := ctx.Value(interceptor.ContextKeyUser).(int64)
	cards, err := cq.getter.GetAll(ctx, userID)
	if err != nil {
		return nil, err
	}
	return createGetAllCardsResponse(cards), nil
}

func createGetCardByIDResponse(card *entity.UserCard) *api.GetCardByIDResponse {
	return &api.GetCardByIDResponse{
		Card: createProtoCard(card),
	}
}

func createGetAllCardsResponse(cards []*entity.UserCard) *api.GetAllCardsResponse {
	resp := &api.GetAllCardsResponse{}
	for _, card := range cards {
		resp.Cards = append(resp.Cards, createProtoCard(card))
	}
	return resp
}

func createProtoCard(card *entity.UserCard) *api.Card {
	return &api.Card{
		Id:           hashids.ID(card.ID).EncodeString(),
		UserId:       hashids.ID(card.UserID).EncodeString(),
		WalletId:     hashids.ID(card.WalletID).EncodeString(),
		LimitDaily:   card.LimitDaily,
		LimitMonthly: card.LimitMonthly,
		CreatedAt:    timestamppb.New(card.CreatedAt),
		UpdatedAt:    timestamppb.New(card.UpdatedAt),
		DeletedAt:    nil,
	}
}
