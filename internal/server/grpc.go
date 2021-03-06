package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logsettable "github.com/grpc-ecosystem/go-grpc-middleware/logging/settable"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/indrasaputra/spenmo/internal/config"
	"github.com/indrasaputra/spenmo/internal/grpc/interceptor"
)

const (
	connProtocol = "tcp"
)

// Closer is responsible to close any open or available resource.
type Closer func()

// Grpc is responsible to act as gRPC server.
// It composes grpc.Server.
type Grpc struct {
	*grpc.Server
	listener net.Listener
	port     string
}

// newGrpc creates an instance of Grpc.
func newGrpc(port string, options ...grpc.ServerOption) *Grpc {
	srv := grpc.NewServer(options...)
	return &Grpc{
		Server: srv,
		port:   port,
	}
}

// NewGrpc creates an instance of Grpc for used in development environment.
//
// These are list of interceptors that are attached (from innermost to outermost):
// 	- Metrics, using Prometheus.
// 	- Logging, using zap logger.
// 	- Recoverer, using grpc_recovery.
func NewGrpc(port string, rate *config.RateLimit) *Grpc {
	options := grpc_middleware.WithUnaryServerChain(defaultUnaryServerInterceptors(rate)...)
	srv := newGrpc(port, options)
	grpc_prometheus.Register(srv.Server)
	return srv
}

// Run runs the server.
// It basically runs grpc.Server.Serve and is a blocking.
func (g *Grpc) Run() error {
	var err error
	g.listener, err = net.Listen(connProtocol, fmt.Sprintf(":%s", g.port))
	if err != nil {
		return err
	}

	go g.serve()
	log.Printf("grpc server is running on port %s\n", g.port)
	return nil
}

// AwaitTermination blocks the server and wait for termination signal.
// The termination signal must be one of SIGINT or SIGTERM.
// Once it receives one of those signals, the gRPC server will perform graceful stop and close the listener.
//
// It receives Closer and will perform all closers before closing itself.
func (g *Grpc) AwaitTermination(closer Closer) error {
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	<-sign

	closer()
	g.GracefulStop()
	return g.listener.Close()
}

func (g *Grpc) serve() {
	if err := g.Serve(g.listener); err != nil {
		panic(err)
	}
}

func defaultUnaryServerInterceptors(rate *config.RateLimit) []grpc.UnaryServerInterceptor {
	logger, _ := zap.NewProduction() // error is impossible, hence ignored.
	grpc_zap.SetGrpcLoggerV2(grpc_logsettable.ReplaceGrpcLoggerV2(), logger)
	grpc_prometheus.EnableHandlingTimeHistogram()

	options := []grpc.UnaryServerInterceptor{
		grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(recoveryHandler)),
		grpc_zap.UnaryServerInterceptor(logger),
		grpc_prometheus.UnaryServerInterceptor,
		otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer()),
		interceptor.OpenTracingUnaryServerInterceptor(),
		interceptor.AuthUnaryServerInterceptor(),
		interceptor.RateLimitUnaryServerInterceptor(rate.RatePerSecond, rate.BurstPerSecond),
	}
	return options
}

func recoveryHandler(p interface{}) error {
	return status.Errorf(codes.Unknown, "%v", p)
}
