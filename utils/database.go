package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	_ "github.com/lib/pq"
)

var Database *pgxpool.Pool

func Connect() error {
	var err error
	Database, err = pgxpool.Connect(context.Background(), GetEnvVariable("POSTGRES_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return nil
}
