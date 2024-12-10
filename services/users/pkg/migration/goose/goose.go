package goose

import (
	"context"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
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

	fmt.Println("dir", dir)

	// FIXME IDK why dir formats from ../../migrations to ./migrations. ?????
	return goose.RunContext(ctx, command, db, dir, []string{dbString, command}...)
}
