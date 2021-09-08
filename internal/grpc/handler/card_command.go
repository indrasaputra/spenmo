package handler

import (
	"context"

	"github.com/indrasaputra/hashids"

	"github.com/indrasaputra/spenmo/entity"
	"github.com/indrasaputra/spenmo/internal/grpc/interceptor"
	api "github.com/indrasaputra/spenmo/proto/indrasaputra/spenmo/v1"
	"github.com/indrasaputra/spenmo/service"
)

// CardCommand handles HTTP/2 gRPC request for state-changing toggle .
type CardCommand struct {
	api.UnimplementedCardCommandServiceServer

	creator service.CreateCard
	updater service.UpdateCard
}

// NewCardCommand creates an instance of CardCommand.
func NewCardCommand(creator service.CreateCard, updater service.UpdateCard) *CardCommand {
	return &CardCommand{
		creator: creator,
		updater: updater,
	}
}

// CreateCard handles HTTP/2 gRPC request similar to POST in HTTP/1.1.
func (cc *CardCommand) CreateCard(ctx context.Context, request *api.CreateCardRequest) (*api.CreateCardResponse, error) {
	if request == nil || request.GetCard() == nil {
		return nil, entity.ErrEmptyCard()
	}

	userID := ctx.Value(interceptor.ContextKeyUser).(int64)
	walletID, err := hashids.DecodeHash([]byte(request.GetCard().GetWalletId()))
	if err != nil {
		return nil, entity.ErrInvalidWallet()
	}

	err = cc.creator.Create(ctx, createCardFromCreateCardRequest(request, userID, int64(walletID)))
	if err != nil {
		return nil, err
	}
	return &api.CreateCardResponse{}, nil
}

// UpdateCard handles HTTP/2 gRPC request similar to PUT in HTTP/1.1.
func (cc *CardCommand) UpdateCard(ctx context.Context, request *api.UpdateCardRequest) (*api.UpdateCardResponse, error) {
	if request == nil || request.GetCard() == nil {
		return nil, entity.ErrEmptyCard()
	}

	userID := ctx.Value(interceptor.ContextKeyUser).(int64)
	cardID, err := hashids.DecodeHash([]byte(request.GetId()))
	if err != nil {
		return nil, entity.ErrInvalidID()
	}

	err = cc.updater.Update(ctx, createCardFromUpdateCardRequest(request, userID, int64(cardID)))
	if err != nil {
		return nil, err
	}
	return &api.UpdateCardResponse{}, nil
}

// DeleteCard handles HTTP/2 gRPC request similar to DELETE in HTTP/1.1.
func (cc *CardCommand) DeleteCard(ctx context.Context, request *api.DeleteCardRequest) (*api.DeleteCardResponse, error) {
	return &api.DeleteCardResponse{}, nil
}

func createCardFromCreateCardRequest(request *api.CreateCardRequest, userID, walletID int64) *entity.UserCard {
	return &entity.UserCard{
		UserID:       userID,
		WalletID:     walletID,
		LimitDaily:   request.GetCard().GetLimitDaily(),
		LimitMonthly: request.GetCard().GetLimitMonthly(),
	}
}

func createCardFromUpdateCardRequest(request *api.UpdateCardRequest, userID, cardID int64) *entity.UserCard {
	return &entity.UserCard{
		ID:           cardID,
		UserID:       userID,
		LimitDaily:   request.GetCard().GetLimitDaily(),
		LimitMonthly: request.GetCard().GetLimitMonthly(),
	}
}
