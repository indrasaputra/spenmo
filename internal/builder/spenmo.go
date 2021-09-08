package builder

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/indrasaputra/spenmo/internal/config"
	"github.com/indrasaputra/spenmo/internal/grpc/handler"
	"github.com/indrasaputra/spenmo/internal/repository/postgres"
	"github.com/indrasaputra/spenmo/service"
)

// BuildCardCommandHandler builds card command handler including all of its dependencies.
func BuildCardCommandHandler(pool *pgxpool.Pool) *handler.CardCommand {
	psql := postgres.NewCard(pool)
	creator := service.NewCardCreator(psql)
	return handler.NewCardCommand(creator)
}

// BuildCardQueryHandler builds card command handler including all of its dependencies.
func BuildCardQueryHandler(pool *pgxpool.Pool) *handler.CardQuery {
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
