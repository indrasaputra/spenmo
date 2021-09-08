package interceptor_test

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	"github.com/indrasaputra/spenmo/internal/grpc/interceptor"
	api "github.com/indrasaputra/spenmo/proto/indrasaputra/spenmo/v1"
	mock_server "github.com/indrasaputra/spenmo/test/mock/proto/grpc/server"
)

const (
	buffer = 1024 * 1024
)

var (
	testCtxNoAuth    = context.Background()
	testCtxWrongAuth = context.WithValue(context.Background(), interceptor.ContextKeyUser, "abc")
	testCtx          = metadata.AppendToOutgoingContext(testCtxNoAuth, "authorization", "1")
)

type CardClientExecutor struct {
	client api.CardCommandServiceClient
	closer func()
}

func TestAuthUnaryServerInterceptor(t *testing.T) {
	t.Run("authorization doesn't exist", func(t *testing.T) {
		exec := createClientExecutor(interceptor.AuthUnaryServerInterceptor())
		defer exec.closer()

		resp, err := exec.client.CreateCard(testCtxNoAuth, &api.CreateCardRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, codes.Unauthenticated, status.Code(err))
		assert.Nil(t, resp)
	})

	t.Run("authorization is not integer", func(t *testing.T) {
		exec := createClientExecutor(interceptor.AuthUnaryServerInterceptor())
		defer exec.closer()

		resp, err := exec.client.CreateCard(testCtxWrongAuth, &api.CreateCardRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, codes.Unauthenticated, status.Code(err))
		assert.Nil(t, resp)
	})

	t.Run("success authenticate", func(t *testing.T) {
		exec := createClientExecutor(interceptor.AuthUnaryServerInterceptor())
		defer exec.closer()

		resp, err := exec.client.CreateCard(testCtx, &api.CreateCardRequest{})

		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
}

func TestOpenTracingUnaryServerInterceptor(t *testing.T) {
	t.Run("success create a new span and finish", func(t *testing.T) {
		exec := createClientExecutor(interceptor.OpenTracingUnaryServerInterceptor())
		defer exec.closer()

		resp, err := exec.client.CreateCard(testCtx, &api.CreateCardRequest{})

		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
}

func createClientExecutor(intercept grpc.UnaryServerInterceptor) *CardClientExecutor {
	listener := bufconn.Listen(buffer)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(intercept))
	spenmoServer := &mock_server.MockCardServiceServer{}
	api.RegisterCardCommandServiceServer(grpcServer, spenmoServer)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			panic(err)
		}
	}()

	conn, err := grpc.DialContext(context.Background(), "", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}), grpc.WithInsecure())
	if err != nil {
		grpcServer.Stop()
		panic(err)
	}

	closer := func() {
		_ = listener.Close()
		grpcServer.GracefulStop()
	}

	return &CardClientExecutor{
		client: api.NewCardCommandServiceClient(conn),
		closer: closer,
	}
}
