package builder_test

import (
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/spenmo/internal/builder"
	"github.com/indrasaputra/spenmo/internal/config"
)

func TestBuildCardCommandHandler(t *testing.T) {
	t.Run("success create card command handler", func(t *testing.T) {
		psql := &pgxpool.Pool{}
		handler := builder.BuildCardCommandHandler(psql)
		assert.NotNil(t, handler)
	})
}

func TestBuildCardQueryHandler(t *testing.T) {
	t.Run("success create card query handler", func(t *testing.T) {
		psql := &pgxpool.Pool{}
		handler := builder.BuildCardQueryHandler(psql)
		assert.NotNil(t, handler)
	})
}

func TestBuildPgxPool(t *testing.T) {
	cfg := &config.Postgres{
		Host:            "localhost",
		Port:            "5432",
		Name:            "spenmo",
		User:            "user",
		Password:        "password",
		MaxOpenConns:    "10",
		MaxConnLifetime: "10m",
		MaxIdleLifetime: "5m",
	}

	t.Run("fail build sql client", func(t *testing.T) {
		client, err := builder.BuildPgxPool(cfg)

		assert.NotNil(t, err)
		assert.Nil(t, client)
	})
}
