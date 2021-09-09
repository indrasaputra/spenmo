package interceptor

import (
	"context"
	"path"
	"strconv"

	"github.com/opentracing/opentracing-go"
	"golang.org/x/time/rate"
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

// RateLimitUnaryServerInterceptor intercepts the request and check if the request is allowed to go through.
func RateLimitUnaryServerInterceptor(ratePerSecond, burst int) grpc.UnaryServerInterceptor {
	limiters := make(map[int64]*rate.Limiter)

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		userID, err := getUserFromContext(ctx)
		if err != nil {
			return nil, status.Error(codes.ResourceExhausted, "unrecognized user has no quota")
		}

		lim, ok := limiters[userID]
		if !ok {
			lim = rate.NewLimiter(rate.Limit(ratePerSecond), burst)
			limiters[userID] = lim
		}
		if !lim.Allow() {
			return nil, status.Error(codes.ResourceExhausted, "quota is exhausted")
		}
		return handler(ctx, req)
	}
}

// AuthUnaryServerInterceptor intercepts the request and check for authentication header.
func AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		id, err := getUserFromContext(ctx)
		if err != nil {
			return nil, err
		}
		return handler(context.WithValue(ctx, ContextKeyUser, id), req)
	}
}

func getUserFromContext(ctx context.Context) (int64, error) {
	userID, ok := ctx.Value(ContextKeyUser).(int64)
	if ok {
		return userID, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, status.Error(codes.Unauthenticated, "request is unauthenticated")
	}

	val := md[authKey]
	if len(val) == 0 {
		return 0, status.Error(codes.Unauthenticated, "request is unauthenticated")
	}

	id, err := strconv.ParseInt(val[0], 10, 64)
	if err != nil {
		return 0, status.Error(codes.Unauthenticated, "request is unauthenticated")
	}
	return id, nil
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
