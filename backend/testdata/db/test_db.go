package db

import (
	"context"
	"log"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var pgContainer *postgres.PostgresContainer

// StartTestDatabase initializes a postgres testcontainer
func StartTestDatabase(ctx context.Context) error {
	var err error
	pgContainer, err = postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithInitScripts("../testdata/db/init-db.sql"),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return err
	}
	return nil
}

// GetTestDatabaseConnection return a connection string that can be used to connect to the test db
func GetTestDatabaseConnection(ctx context.Context) (string, error) {
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Printf("failed to get test postgres container connection string: %s", err)
		return "", err
	}
	return connStr, nil
}

// StopTestDatabase closes the postgres testcontainer
func StopTestDatabase() {
	if err := testcontainers.TerminateContainer(pgContainer); err != nil {
		log.Printf("failed to terminate container: %s", err)
	}
}
