package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

// repository errors definitions
var (
	ErrNotFound        = errors.New("not found")
	ErrAlreadyExists   = errors.New("already exists")
	ErrInternalFailure = errors.New("internal failure") // FIXME пока-что не юзаю, а смысл?
)

func NewPool(dsn string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		panic(err)
	}

	if err = pool.Ping(context.Background()); err != nil {
		panic(err)
	}

	return pool
}
