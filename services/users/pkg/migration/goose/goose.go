package goose

import (
	"context"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log"
)

func Run(ctx context.Context, dir, dbString, command string) error {
	db, err := goose.OpenDBWithDriver("postgres", dbString)
	if err != nil {
		return err
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	return goose.RunContext(ctx, command, db, dir, []string{dbString, command}...)
}
