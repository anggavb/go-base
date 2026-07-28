package config_postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(dbDSN string) (*pgxpool.Pool, error) {
	pg, _ := pgxpool.New(context.Background(), dbDSN)
	return pg, pg.Ping(context.Background())
}
