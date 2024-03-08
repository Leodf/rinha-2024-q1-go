package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var PG *pgxpool.Pool

var (
	dbconfig = map[string]string{
		"DB_HOST":     "localhost",
		"DB_PORT":     "5432",
		"DB_USER":     "admin",
		"DB_PASSWORD": "123",
		"DB_NAME":     "rinha",
		"DB_POOL_MAX": "1",
		"DB_POOL_MIN": "1",
	}
)

func Init() (*pgxpool.Pool, error) {
	ctx := context.Background()
	for key := range dbconfig {
		if v, ok := os.LookupEnv(key); ok {
			dbconfig[key] = v
		}
	}
	uri := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable pool_min_conns=%s pool_max_conns=%s",
		dbconfig["DB_HOST"],
		dbconfig["DB_PORT"],
		dbconfig["DB_USER"],
		dbconfig["DB_PASSWORD"],
		dbconfig["DB_NAME"],
		dbconfig["DB_POOL_MAX"],
		dbconfig["DB_POOL_MIN"],
	)
	cfg, err := pgxpool.ParseConfig(uri)
	if err != nil {
		return nil, fmt.Errorf("cannot parse config from conn string: %w", err)
	}

	db, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	PG = db
	return db, nil
}
