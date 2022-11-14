package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

var Pool *pgxpool.Pool

// ConnectToPostgres Allow the connection to the database.
func ConnectToPostgres() {
	var err error
	//dsn := os.Getenv("DSN")
	Pool, err = pgxpool.New(context.Background(), "postgres://postgres:postgres@127.0.0.1:5432/ulule")
	//Pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}
