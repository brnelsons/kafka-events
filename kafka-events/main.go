package main

import (
	"fmt"
	"kafka-events/postgres"
	"os"
	"time"
)

const (
	EnvDbHost           = "DB_HOST"
	EnvDbPort           = "DB_PORT"
	EnvDbSchema         = "DB_SCHEMA"
	EnvDbUser           = "DB_USER"
	EnvDbPassword       = "DB_PASSWORD"
	EnvDbConnectTimeout = "DB_CONNECT_TIMEOUT"
)

func main() {
	dbHost := os.Getenv(EnvDbHost)
	dbPort := os.Getenv(EnvDbPort)
	dbSchema := os.Getenv(EnvDbSchema)
	dbUser := os.Getenv(EnvDbUser)
	dbPassword := os.Getenv(EnvDbPassword)
	dbConnectTimeout, err := time.ParseDuration(orElse(os.Getenv(EnvDbConnectTimeout), "10s"))

	postgresConnection := postgres.New(dbHost, dbPort, dbSchema, dbUser, dbPassword)

	err = postgresConnection.Connect(dbConnectTimeout)
	if err != nil {
		panic(err)
	}
	err = postgresConnection.Migrate()
	if err != nil {
		panic(err)
	}

	// TODO start webservices (Full CRUD operations)
	// TODO start kafka publishers (Just publishes event notifications)
	// upon receiving a create request for an event via REST endpoints
	// 1. process the event
	// 2. save the event to the database
	// 3. publish event to kafka stream (created stream)

	fmt.Println("We're up!")
}

func orElse(this string, that string) string {
	if this != "" {
		return this
	}
	return that
}
