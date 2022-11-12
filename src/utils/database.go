package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

var Pool *pgxpool.Pool

func ConnectToPostgres() {
	var err error
	Pool, err = pgxpool.New(context.Background(), "postgres://postgres:postgres@127.0.0.1:5432/ulule")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}
