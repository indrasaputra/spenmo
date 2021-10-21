package builder_test

import (
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/spenmo/internal/builder"
	"github.com/indrasaputra/spenmo/internal/config"
	"github.com/indrasaputra/spenmo/internal/repository/model/ent"
)

func TestBuildCardCommandHandlerUsingPgx(t *testing.T) {
	t.Run("success create card command handler", func(t *testing.T) {
		psql := &pgxpool.Pool{}
		handler := builder.BuildCardCommandHandlerUsingPgx(psql)
		assert.NotNil(t, handler)
	})
}

func TestBuildCardQueryHandlerUsingPgx(t *testing.T) {
	t.Run("success create card query handler", func(t *testing.T) {
		psql := &pgxpool.Pool{}
		handler := builder.BuildCardQueryHandlerUsingPgx(psql)
		assert.NotNil(t, handler)
	})
}

func TestBuildCardCommandHandlerUsingEnt(t *testing.T) {
	t.Run("success create card command handler using ent", func(t *testing.T) {
		client := &ent.Client{}
		handler := builder.BuildCardCommandHandlerUsingEnt(client)
		assert.NotNil(t, handler)
	})
}

func TestBuildCardQueryHandlerUsingEnt(t *testing.T) {
	t.Run("success create card query handler using ent", func(t *testing.T) {
		client := &ent.Client{}
		handler := builder.BuildCardQueryHandlerUsingEnt(client)
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

func TestBuildEntPgxClient(t *testing.T) {
	cfg := &config.Postgres{
		Host:     "localhost",
		Port:     "5432",
		Name:     "spenmo",
		User:     "user",
		Password: "password",
	}

	t.Run("success build ent pgx client", func(t *testing.T) {
		client, err := builder.BuildEntPgxClient(cfg)
		defer func() { _ = client.Close() }()

		assert.Nil(t, err)
		assert.NotNil(t, client)
	})
}
