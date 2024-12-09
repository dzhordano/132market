package postgres

import (
	"context"
	"errors"

	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrAlreadyExists   = errors.New("already exists")
	ErrInternalFailure = errors.New("internal failure") // FIXME пока-что не юзаю, а смысл?
)

func NewPool(dsn string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v. DSN: %s", err, dsn)
	}

	if err = pool.Ping(context.Background()); err != nil {
		log.Fatalf("Error pinging database: %v. DSN: %s", err, dsn)
	}

	return pool
}
