package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// Config holds configuration for the project.
type Config struct {
	ServiceName string `env:"SERVICE_NAME,default=spenmo-api"`
	Port        Port
	Postgres    Postgres
	Jaeger      Jaeger
	RateLimit   RateLimit
}

// Port holds configuration for project's port.
type Port struct {
	GRPC string `env:"PORT_GRPC,default=8080"`
	REST string `env:"PORT_REST,default=8081"`
}

// Postgres holds all configuration for PostgreSQL.
type Postgres struct {
	Host            string `env:"POSTGRES_HOST,default=localhost"`
	Port            string `env:"POSTGRES_PORT,default=5432"`
	User            string `env:"POSTGRES_USER,required"`
	Password        string `env:"POSTGRES_PASSWORD,required"`
	Name            string `env:"POSTGRES_NAME,required"`
	MaxOpenConns    string `env:"POSTGRES_MAX_OPEN_CONNS,default=5"`
	MaxConnLifetime string `env:"POSTGRES_MAX_CONN_LIFETIME,default=10m"`
	MaxIdleLifetime string `env:"POSTGRES_MAX_IDLE_LIFETIME,default=5m"`
}

// Jaeger holds configuration for Jaeger.
type Jaeger struct {
	Enabled       bool    `env:"JAEGER_ENABLED,default=true"`
	Host          string  `env:"JAEGER_HOST,default=localhost"`
	Port          string  `env:"JAEGER_PORT,default=6831"`
	SamplingType  string  `env:"JAEGER_SAMPLING_TYPE,default=const"`
	SamplingParam float64 `env:"JAEGER_SAMPLING_PARAM,default=1"`
	LogSpans      bool    `env:"JAEGER_LOG_SPANS,default=true"`
	FlushInterval uint    `env:"JAEGER_FLUSH_INTERVAL,default=1"`
}

// RateLimit holds configuration for RateLimit.
type RateLimit struct {
	RatePerSecond  int `env:"RATE_LIMIT_PER_SECOND,default=1"`
	BurstPerSecond int `env:"RATE_BURST_PER_SECOND,default=1"`
}

// NewConfig creates an instance of Config.
// It needs the path of the env file to be used.
func NewConfig(env string) (*Config, error) {
	// just skip loading env files if it is not exists, env files only used in local dev
	_ = godotenv.Load(env)

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding env")
	}

	return &config, nil
}
