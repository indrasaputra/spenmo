package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/indrasaputra/spenmo/entity"
	"github.com/indrasaputra/spenmo/internal/builder"
	"github.com/indrasaputra/spenmo/internal/config"
)

var (
	psql    *pgxpool.Pool
	errPsql error
	ctx     = context.Background()
)

func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)

	psql, errPsql = builder.BuildPgxPool(&cfg.Postgres)
	checkError(errPsql)

	users := []*entity.User{
		{ID: 1, Name: "a", Email: "a@a"},
		{ID: 2, Name: "b", Email: "b@b"},
		{ID: 3, Name: "c", Email: "c@c"},
		{ID: 4, Name: "d", Email: "d@d"},
		{ID: 5, Name: "e", Email: "e@e"},
	}

	numWalletPerUser := 5
	walletBalance := float64(1000000)
	for _, user := range users {
		if err := insertUser(user); err != nil {
			log.Println(err)
			continue
		}
		for i := 0; i < numWalletPerUser; i++ {
			if err := insertUserWallet(user.ID, walletBalance); err != nil {
				log.Println(err)
			}
		}
	}
}

func insertUser(user *entity.User) error {
	query := "INSERT INTO " +
		"users (name, email, password, created_at, updated_at) " +
		"VALUES ($1, $2, $3, $4, $5)"

	_, err := psql.Exec(ctx, query,
		user.Name,
		user.Email,
		user.Email,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	return err
}

func insertUserWallet(userID int64, balance float64) error {
	query := "INSERT INTO " +
		"user_wallets (user_id, balance, created_at, updated_at) " +
		"VALUES ($1, $2, $3, $4)"

	_, err := psql.Exec(ctx, query,
		userID,
		balance,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	return err
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
