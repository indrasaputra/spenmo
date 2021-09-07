// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package spenmov1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CardCommandServiceClient is the client API for CardCommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardCommandServiceClient interface {
	// Create a new card.
	//
	// This endpoint creates a new card with provided walletId, limitDaily, and
	// limitMonthly. The user information is taken from authentication.
	CreateCard(ctx context.Context, in *CreateCardRequest, opts ...grpc.CallOption) (*CreateCardResponse, error)
	// Update an existing card.
	//
	// This endpoint update card's information.
	// The user information is taken from authentication.
	UpdateCard(ctx context.Context, in *UpdateCardRequest, opts ...grpc.CallOption) (*UpdateCardResponse, error)
	// Delete an existing card.
	//
	// This endpoint deletes a card by its id.
	// The operation is soft-delete, thus the card will stay in storage.
	// The user information is taken from authentication.
	DeleteCard(ctx context.Context, in *DeleteCardRequest, opts ...grpc.CallOption) (*DeleteCardResponse, error)
}

type cardCommandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCardCommandServiceClient(cc grpc.ClientConnInterface) CardCommandServiceClient {
	return &cardCommandServiceClient{cc}
}

func (c *cardCommandServiceClient) CreateCard(ctx context.Context, in *CreateCardRequest, opts ...grpc.CallOption) (*CreateCardResponse, error) {
	out := new(CreateCardResponse)
	err := c.cc.Invoke(ctx, "/proto.indrasaputra.spenmo.v1.CardCommandService/CreateCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardCommandServiceClient) UpdateCard(ctx context.Context, in *UpdateCardRequest, opts ...grpc.CallOption) (*UpdateCardResponse, error) {
	out := new(UpdateCardResponse)
	err := c.cc.Invoke(ctx, "/proto.indrasaputra.spenmo.v1.CardCommandService/UpdateCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardCommandServiceClient) DeleteCard(ctx context.Context, in *DeleteCardRequest, opts ...grpc.CallOption) (*DeleteCardResponse, error) {
	out := new(DeleteCardResponse)
	err := c.cc.Invoke(ctx, "/proto.indrasaputra.spenmo.v1.CardCommandService/DeleteCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardCommandServiceServer is the server API for CardCommandService service.
// All implementations must embed UnimplementedCardCommandServiceServer
// for forward compatibility
type CardCommandServiceServer interface {
	// Create a new card.
	//
	// This endpoint creates a new card with provided walletId, limitDaily, and
	// limitMonthly. The user information is taken from authentication.
	CreateCard(context.Context, *CreateCardRequest) (*CreateCardResponse, error)
	// Update an existing card.
	//
	// This endpoint update card's information.
	// The user information is taken from authentication.
	UpdateCard(context.Context, *UpdateCardRequest) (*UpdateCardResponse, error)
	// Delete an existing card.
	//
	// This endpoint deletes a card by its id.
	// The operation is soft-delete, thus the card will stay in storage.
	// The user information is taken from authentication.
	DeleteCard(context.Context, *DeleteCardRequest) (*DeleteCardResponse, error)
	mustEmbedUnimplementedCardCommandServiceServer()
}

// UnimplementedCardCommandServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCardCommandServiceServer struct {
}

func (UnimplementedCardCommandServiceServer) CreateCard(context.Context, *CreateCardRequest) (*CreateCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCard not implemented")
}
func (UnimplementedCardCommandServiceServer) UpdateCard(context.Context, *UpdateCardRequest) (*UpdateCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCard not implemented")
}
func (UnimplementedCardCommandServiceServer) DeleteCard(context.Context, *DeleteCardRequest) (*DeleteCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCard not implemented")
}
func (UnimplementedCardCommandServiceServer) mustEmbedUnimplementedCardCommandServiceServer() {}

// UnsafeCardCommandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardCommandServiceServer will
// result in compilation errors.
type UnsafeCardCommandServiceServer interface {
	mustEmbedUnimplementedCardCommandServiceServer()
}

func RegisterCardCommandServiceServer(s grpc.ServiceRegistrar, srv CardCommandServiceServer) {
	s.RegisterService(&CardCommandService_ServiceDesc, srv)
}

func _CardCommandService_CreateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardCommandServiceServer).CreateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.indrasaputra.spenmo.v1.CardCommandService/CreateCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardCommandServiceServer).CreateCard(ctx, req.(*CreateCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardCommandService_UpdateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardCommandServiceServer).UpdateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.indrasaputra.spenmo.v1.CardCommandService/UpdateCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardCommandServiceServer).UpdateCard(ctx, req.(*UpdateCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardCommandService_DeleteCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardCommandServiceServer).DeleteCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.indrasaputra.spenmo.v1.CardCommandService/DeleteCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardCommandServiceServer).DeleteCard(ctx, req.(*DeleteCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CardCommandService_ServiceDesc is the grpc.ServiceDesc for CardCommandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CardCommandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.indrasaputra.spenmo.v1.CardCommandService",
	HandlerType: (*CardCommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCard",
			Handler:    _CardCommandService_CreateCard_Handler,
		},
		{
			MethodName: "UpdateCard",
			Handler:    _CardCommandService_UpdateCard_Handler,
		},
		{
			MethodName: "DeleteCard",
			Handler:    _CardCommandService_DeleteCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/indrasaputra/spenmo/v1/spenmo.proto",
}

// CardQueryServiceClient is the client API for CardQueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardQueryServiceClient interface {
	// Get a card information.
	//
	// This endpoint gets a single card by its key.
	// The user information is taken from authentication.
	GetCardByID(ctx context.Context, in *GetCardByIDRequest, opts ...grpc.CallOption) (*GetCardByIDResponse, error)
	// Get all user's cards.
	//
	// This endpoint gets all available user's cards in the system.
	GetAllCards(ctx context.Context, in *GetAllCardsRequest, opts ...grpc.CallOption) (*GetAllCardsResponse, error)
}

type cardQueryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCardQueryServiceClient(cc grpc.ClientConnInterface) CardQueryServiceClient {
	return &cardQueryServiceClient{cc}
}

func (c *cardQueryServiceClient) GetCardByID(ctx context.Context, in *GetCardByIDRequest, opts ...grpc.CallOption) (*GetCardByIDResponse, error) {
	out := new(GetCardByIDResponse)
	err := c.cc.Invoke(ctx, "/proto.indrasaputra.spenmo.v1.CardQueryService/GetCardByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardQueryServiceClient) GetAllCards(ctx context.Context, in *GetAllCardsRequest, opts ...grpc.CallOption) (*GetAllCardsResponse, error) {
	out := new(GetAllCardsResponse)
	err := c.cc.Invoke(ctx, "/proto.indrasaputra.spenmo.v1.CardQueryService/GetAllCards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardQueryServiceServer is the server API for CardQueryService service.
// All implementations must embed UnimplementedCardQueryServiceServer
// for forward compatibility
type CardQueryServiceServer interface {
	// Get a card information.
	//
	// This endpoint gets a single card by its key.
	// The user information is taken from authentication.
	GetCardByID(context.Context, *GetCardByIDRequest) (*GetCardByIDResponse, error)
	// Get all user's cards.
	//
	// This endpoint gets all available user's cards in the system.
	GetAllCards(context.Context, *GetAllCardsRequest) (*GetAllCardsResponse, error)
	mustEmbedUnimplementedCardQueryServiceServer()
}

// UnimplementedCardQueryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCardQueryServiceServer struct {
}

func (UnimplementedCardQueryServiceServer) GetCardByID(context.Context, *GetCardByIDRequest) (*GetCardByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCardByID not implemented")
}
func (UnimplementedCardQueryServiceServer) GetAllCards(context.Context, *GetAllCardsRequest) (*GetAllCardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCards not implemented")
}
func (UnimplementedCardQueryServiceServer) mustEmbedUnimplementedCardQueryServiceServer() {}

// UnsafeCardQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardQueryServiceServer will
// result in compilation errors.
type UnsafeCardQueryServiceServer interface {
	mustEmbedUnimplementedCardQueryServiceServer()
}

func RegisterCardQueryServiceServer(s grpc.ServiceRegistrar, srv CardQueryServiceServer) {
	s.RegisterService(&CardQueryService_ServiceDesc, srv)
}

func _CardQueryService_GetCardByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardQueryServiceServer).GetCardByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.indrasaputra.spenmo.v1.CardQueryService/GetCardByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardQueryServiceServer).GetCardByID(ctx, req.(*GetCardByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardQueryService_GetAllCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardQueryServiceServer).GetAllCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.indrasaputra.spenmo.v1.CardQueryService/GetAllCards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardQueryServiceServer).GetAllCards(ctx, req.(*GetAllCardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CardQueryService_ServiceDesc is the grpc.ServiceDesc for CardQueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CardQueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.indrasaputra.spenmo.v1.CardQueryService",
	HandlerType: (*CardQueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCardByID",
			Handler:    _CardQueryService_GetCardByID_Handler,
		},
		{
			MethodName: "GetAllCards",
			Handler:    _CardQueryService_GetAllCards_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/indrasaputra/spenmo/v1/spenmo.proto",
}
