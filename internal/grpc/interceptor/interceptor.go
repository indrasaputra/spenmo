package interceptor

import (
	"context"
	"path"
	"strconv"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// ContextKey is just an alias for string to be used
// as key when assign a value in context.
// This is to avoid go-lint warning
// "should not use basic type untyped string as key in context.WithValue".
type ContextKey string

const (
	// ContextKeyUser is just a string "user" defined as a key
	// to save a user information in context.
	ContextKeyUser = ContextKey("user")

	tagMethod = "grpc.method"
	tagCode   = "grpc.code"
	authKey   = "authorization"
)

// AuthUnaryServerInterceptor intercepts the request and check for authentication header.
func AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "request is unauthenticated")
		}

		val := md[authKey]
		if len(val) == 0 {
			return nil, status.Error(codes.Unauthenticated, "request is unauthenticated")
		}

		id, err := strconv.ParseInt(val[0], 10, 64)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "request is unauthenticated")
		}
		return handler(context.WithValue(ctx, ContextKeyUser, id), req)
	}
}

// OpenTracingUnaryServerInterceptor intercepts the request and creates a span from the incoming context.
// It names the span using the method that is being called.
func OpenTracingUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		method := path.Base(info.FullMethod)
		span, ctx := opentracing.StartSpanFromContext(ctx, method)
		defer span.Finish()

		resp, err := handler(ctx, req)

		span.SetTag(tagMethod, method)
		span.SetTag(tagCode, status.Code(err))

		return resp, err
	}
}
