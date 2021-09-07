// DO NOT EDIT.
// Umm.. actually you can edit this file :D
// The first sentence was only to avoid golint due to package name that uses underscore.

package mock_server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	api "github.com/indrasaputra/spenmo/proto/indrasaputra/spenmo/v1"
)

var (
	keyErr      = "has-error"
	keyReturn   = "complete-return"
	errInternal = status.New(codes.Internal, "").Err()
)

// MockCardServiceServer must be embedded to have forward compatible implementations.
type MockCardServiceServer struct {
	api.UnimplementedCardCommandServiceServer
	api.UnimplementedCardQueryServiceServer
}

func (MockCardServiceServer) CreateCard(ctx context.Context, _ *api.CreateCardRequest) (*api.CreateCardResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md[keyErr]) > 0 && md[keyErr][0] != "" {
		return nil, errInternal
	}
	return &api.CreateCardResponse{}, nil
}

func (MockCardServiceServer) GetCardByID(ctx context.Context, _ *api.GetCardByIDRequest) (*api.GetCardByIDResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md[keyErr]) > 0 && md[keyErr][0] != "" {
		return nil, errInternal
	}
	if len(md[keyReturn]) > 0 && md[keyReturn][0] != "" {
		return &api.GetCardByIDResponse{Card: &api.Card{Id: md[keyReturn][0]}}, nil
	}
	return &api.GetCardByIDResponse{}, nil
}

func (MockCardServiceServer) GetAllCards(ctx context.Context, _ *api.GetAllCardsRequest) (*api.GetAllCardsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md[keyErr]) > 0 && md[keyErr][0] != "" {
		return nil, errInternal
	}
	return &api.GetAllCardsResponse{}, nil
}

func (MockCardServiceServer) UpdateCard(ctx context.Context, _ *api.UpdateCardRequest) (*api.UpdateCardResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md[keyErr]) > 0 && md[keyErr][0] != "" {
		return nil, errInternal
	}
	return &api.UpdateCardResponse{}, nil
}

func (MockCardServiceServer) DeleteCard(ctx context.Context, _ *api.DeleteCardRequest) (*api.DeleteCardResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md[keyErr]) > 0 && md[keyErr][0] != "" {
		return nil, errInternal
	}
	return &api.DeleteCardResponse{}, nil
}
