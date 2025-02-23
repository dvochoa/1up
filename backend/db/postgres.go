package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetDatabaseConnection() string {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Fatalf("Error casting POSTGRES_PORT to int: %v", err)
	}
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
}
