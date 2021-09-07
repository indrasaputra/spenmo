package entity

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/indrasaputra/spenmo/proto/indrasaputra/spenmo/v1"
)

// ErrInternal returns codes.Internal explained that unexpected behavior occurred in system.
func ErrInternal(message string) error {
	st := status.New(codes.Internal, message)
	te := &api.SpenmoCardError{
		ErrorCode: api.SpenmoCardErrorCode_INTERNAL,
	}
	res, err := st.WithDetails(te)
	if err != nil {
		return st.Err()
	}
	return res.Err()
}

// ErrEmptyCard returns codes.InvalidArgument explained that the instance is empty or nil.
func ErrEmptyCard() error {
	st := status.New(codes.InvalidArgument, "")
	br := createBadRequest(&errdetails.BadRequest_FieldViolation{
		Field:       "card instance",
		Description: "empty or nil",
	})

	te := &api.SpenmoCardError{
		ErrorCode: api.SpenmoCardErrorCode_EMPTY_CARD,
	}
	res, err := st.WithDetails(br, te)
	if err != nil {
		return st.Err()
	}
	return res.Err()
}

// ErrInvalidID returns codes.InvalidArgument explained that the instance's id is invalid.
func ErrInvalidID() error {
	st := status.New(codes.InvalidArgument, "")
	br := createBadRequest(&errdetails.BadRequest_FieldViolation{
		Field:       "id",
		Description: "id is not hashid",
	})

	te := &api.SpenmoCardError{
		ErrorCode: api.SpenmoCardErrorCode_INVALID_ID,
	}
	res, err := st.WithDetails(br, te)
	if err != nil {
		return st.Err()
	}
	return res.Err()
}

// ErrNotFound returns codes.NotFound explained that the instance is not found.
func ErrNotFound() error {
	st := status.New(codes.NotFound, "")
	te := &api.SpenmoCardError{
		ErrorCode: api.SpenmoCardErrorCode_NOT_FOUND,
	}
	res, err := st.WithDetails(te)
	if err != nil {
		return st.Err()
	}
	return res.Err()
}

// ErrInvalidUser returns codes.InvalidArgument explained that the instance's user is invalid.
func ErrInvalidUser() error {
	st := status.New(codes.InvalidArgument, "")
	br := createBadRequest(&errdetails.BadRequest_FieldViolation{
		Field:       "userId",
		Description: "user is invalid",
	})

	te := &api.SpenmoCardError{
		ErrorCode: api.SpenmoCardErrorCode_INVALID_USER,
	}
	res, err := st.WithDetails(br, te)
	if err != nil {
		return st.Err()
	}
	return res.Err()
}

// ErrInvalidWallet returns codes.InvalidArgument explained that the instance's wallet is invalid.
func ErrInvalidWallet() error {
	st := status.New(codes.InvalidArgument, "")
	br := createBadRequest(&errdetails.BadRequest_FieldViolation{
		Field:       "walletId",
		Description: "wallet is invalid",
	})

	te := &api.SpenmoCardError{
		ErrorCode: api.SpenmoCardErrorCode_INVALID_WALLET,
	}
	res, err := st.WithDetails(br, te)
	if err != nil {
		return st.Err()
	}
	return res.Err()
}

// ErrInvalidLimit returns codes.InvalidArgument explained that the instance's limit is invalid.
func ErrInvalidLimit() error {
	st := status.New(codes.InvalidArgument, "")
	br := createBadRequest(&errdetails.BadRequest_FieldViolation{
		Field:       "limit*",
		Description: "limit is invalid. it must be greater than 0",
	})

	te := &api.SpenmoCardError{
		ErrorCode: api.SpenmoCardErrorCode_INVALID_LIMIT,
	}
	res, err := st.WithDetails(br, te)
	if err != nil {
		return st.Err()
	}
	return res.Err()
}

func createBadRequest(details ...*errdetails.BadRequest_FieldViolation) *errdetails.BadRequest {
	return &errdetails.BadRequest{
		FieldViolations: details,
	}
}
