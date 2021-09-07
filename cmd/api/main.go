package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerconf "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"

	"github.com/indrasaputra/spenmo/internal/builder"
	"github.com/indrasaputra/spenmo/internal/config"
	"github.com/indrasaputra/spenmo/internal/server"
	api "github.com/indrasaputra/spenmo/proto/indrasaputra/spenmo/v1"
)

func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)

	psql, err := builder.BuildPgxPool(&cfg.Postgres)
	checkError(err)
	trc := initTracing(cfg)

	grpcServer := server.NewGrpc(cfg.Port.GRPC)
	registerGrpcHandlers(grpcServer.Server, cfg, psql)

	restServer := server.NewRest(cfg.Port.REST)
	registerRestHandlers(context.Background(), restServer.ServeMux, fmt.Sprintf(":%s", cfg.Port.GRPC), grpc.WithInsecure())

	closer := func() {
		_ = trc.Close()
		psql.Close()
	}

	_ = grpcServer.Run()
	_ = restServer.Run()
	_ = grpcServer.AwaitTermination(closer)
}

func registerGrpcHandlers(server *grpc.Server, cfg *config.Config, psql *pgxpool.Pool) {
	// start register all module's gRPC handlers
	command := builder.BuildCardCommandHandler(psql)
	api.RegisterCardCommandServiceServer(server, command)
	// end of register all module's gRPC handlers
}

func registerRestHandlers(ctx context.Context, server *runtime.ServeMux, grpcPort string, options ...grpc.DialOption) {
	// start register all module's REST handlers
	err := api.RegisterCardCommandServiceHandlerFromEndpoint(ctx, server, grpcPort, options)
	checkError(err)
	// end of register all module's REST handlers
}

func initTracing(cfg *config.Config) io.Closer {
	if !cfg.Jaeger.Enabled {
		return nopCloser{}
	}

	jaegerCfg := &jaegerconf.Configuration{
		ServiceName: cfg.ServiceName,
		Sampler: &jaegerconf.SamplerConfig{
			Type:  cfg.Jaeger.SamplingType,
			Param: cfg.Jaeger.SamplingParam,
		},
		Reporter: &jaegerconf.ReporterConfig{
			LogSpans:            cfg.Jaeger.LogSpans,
			LocalAgentHostPort:  fmt.Sprintf("%s:%s", cfg.Jaeger.Host, cfg.Jaeger.Port),
			BufferFlushInterval: time.Duration(cfg.Jaeger.FlushInterval) * time.Second,
		},
	}
	tracer, closer, err := jaegerCfg.NewTracer(jaegerconf.Logger(jaeger.StdLogger))
	checkError(err)

	opentracing.SetGlobalTracer(tracer)
	return closer
}

type nopCloser struct{}

// Closer closes nothing.
func (nopCloser) Close() error { return nil }

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
