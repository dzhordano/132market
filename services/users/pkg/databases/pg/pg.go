package pg

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// TODO Посмотреть у Олежи че там по своему бд менеджеру и мб сделать норм клиент.

type DB interface {
	// Exec executes a SQL query with the given arguments.
	// The returned CommandTag contains information about the executed query.
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)

	// Query executes a SQL query with the given arguments and returns the results.
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)

	// QueryRow executes a SQL query with the given arguments and returns a single row.
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

func NewPGPool(ctx context.Context, dsn string) (DB, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
