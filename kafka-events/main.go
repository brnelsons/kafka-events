package main

import (
	"fmt"
	"kafka-events/postgres"
	"os"
)

const (
	EnvDbHost     = "DB_HOST"
	EnvDbPort     = "DB_PORT"
	EnvDbSchema   = "DB_SCHEMA"
	EnvDbUser     = "DB_USER"
	EnvDbPassword = "DB_PASSWORD"
)

func main() {
	dbHost := os.Getenv(EnvDbHost)
	dbPort := os.Getenv(EnvDbPort)
	dbSchema := os.Getenv(EnvDbSchema)
	dbUser := os.Getenv(EnvDbUser)
	dbPassword := os.Getenv(EnvDbPassword)

	postgresConnection := postgres.New(dbHost, dbPort, dbSchema, dbUser, dbPassword)

	err := postgresConnection.Connect()
	if err != nil {
		panic(err)
	}
	err = postgresConnection.Migrate()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success!")
}
