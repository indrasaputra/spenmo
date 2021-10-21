package builder

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/indrasaputra/spenmo/internal/config"
	"github.com/indrasaputra/spenmo/internal/grpc/handler"
	"github.com/indrasaputra/spenmo/internal/repository/model"
	"github.com/indrasaputra/spenmo/internal/repository/model/ent"
	"github.com/indrasaputra/spenmo/internal/repository/postgres"
	"github.com/indrasaputra/spenmo/service"
)

const (
	driverPgx = "pgx"
)

// BuildCardCommandHandlerUsingPgx builds card command handler including all of its dependencies using pgxpool.
func BuildCardCommandHandlerUsingPgx(pool *pgxpool.Pool) *handler.CardCommand {
	psql := postgres.NewCard(pool)

	creator := service.NewCardCreator(psql)
	updater := service.NewCardUpdater(psql)
	deleter := service.NewCardDeleter(psql)

	return handler.NewCardCommand(creator, updater, deleter)
}

// BuildCardQueryHandlerUsingPgx builds card command handler including all of its dependencies using pgxpool.
func BuildCardQueryHandlerUsingPgx(pool *pgxpool.Pool) *handler.CardQuery {
	psql := postgres.NewCard(pool)
	getter := service.NewCardGetter(psql)
	return handler.NewCardQuery(getter)
}

// BuildPgxPool builds a pool of pgx client.
func BuildPgxPool(cfg *config.Postgres) (*pgxpool.Pool, error) {
	connCfg := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable pool_max_conns=%s pool_max_conn_lifetime=%s pool_max_conn_idle_time=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.MaxOpenConns,
		cfg.MaxConnLifetime,
		cfg.MaxIdleLifetime,
	)
	return pgxpool.Connect(context.Background(), connCfg)
}

// BuildCardCommandHandlerUsingEnt builds card command handler including all of its dependencies using ent as ORM.
func BuildCardCommandHandlerUsingEnt(client *ent.Client) *handler.CardCommand {
	cardModel := model.NewCard(client)

	creator := service.NewCardCreator(cardModel)
	updater := service.NewCardUpdater(cardModel)
	deleter := service.NewCardDeleter(cardModel)

	return handler.NewCardCommand(creator, updater, deleter)
}

// BuildCardQueryHandlerUsingEnt builds card command handler including all of its dependencies using ent as ORM.
func BuildCardQueryHandlerUsingEnt(client *ent.Client) *handler.CardQuery {
	cardModel := model.NewCard(client)
	getter := service.NewCardGetter(cardModel)
	return handler.NewCardQuery(getter)
}

// BuildEntPgxClient builds an ent client using pgx as internal implementation.
func BuildEntPgxClient(cfg *config.Postgres) (*ent.Client, error) {
	connCfg := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	db, err := sql.Open(driverPgx, connCfg)
	if err != nil {
		return nil, err
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}
