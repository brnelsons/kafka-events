package main

import (
	"context"
	"kafka-events/kafka"
	"kafka-events/postgres"
	"kafka-events/web"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/cors"
)

const (
	EnvDbHost            = "DB_HOST"
	EnvDbPort            = "DB_PORT"
	EnvDbSchema          = "DB_SCHEMA"
	EnvDbUser            = "DB_USER"
	EnvDbPassword        = "DB_PASSWORD"
	EnvDbConnectTimeout  = "DB_CONNECT_TIMEOUT"
	EnvKafkaPublishTopic = "KAFKA_PUBLISH_TOPIC"
	EnvKafkaBrokerCsv    = "KAFKA_BROKER_CSV"

	ApiV1BasePath = "/api/v1"
)

func main() {
	dbHost := os.Getenv(EnvDbHost)
	dbPort := os.Getenv(EnvDbPort)
	dbSchema := os.Getenv(EnvDbSchema)
	dbUser := os.Getenv(EnvDbUser)
	dbPassword := os.Getenv(EnvDbPassword)
	kafkaPublishTopic := os.Getenv(EnvKafkaPublishTopic)
	kafkaBrokerCsv := os.Getenv(EnvKafkaBrokerCsv)
	dbConnectTimeout, err := time.ParseDuration(orElse(os.Getenv(EnvDbConnectTimeout), "10s"))

	dbService := postgres.NewDbService(dbHost, dbPort, dbSchema, dbUser, dbPassword)

	err = dbService.Connect(dbConnectTimeout)
	if err != nil {
		panic(err)
	}
	err = dbService.Migrate()
	if err != nil {
		panic(err)
	}

	kafkaService := kafka.NewKafkaService(
		kafkaPublishTopic,
		strings.Split(kafkaBrokerCsv, ",")...,
	)

	err = kafkaService.Publish(context.Background(), "testing")
	if err != nil {
		panic(err)
	}

	// upon receiving a create request for an event via REST endpoints
	// 1. process the event
	// 2. save the event to the database
	// 3. publish event to kafka stream (created stream)
	// TODO start webservices (Full CRUD operations)

	router := web.NewRouter(ApiV1BasePath, GetRoutes(kafkaService, dbService))

	log.Println("Successfully started.")
	log.Fatal(http.ListenAndServe(":8081", setupGlobalMiddleware(router)))
}

// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}

func orElse(this string, that string) string {
	if this != "" {
		return this
	}
	return that
}
