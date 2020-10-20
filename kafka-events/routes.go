package main

import (
	"kafka-events/kafka"
	"kafka-events/postgres"
	"kafka-events/web"
	"net/http"
)

func GetRoutes(kafkaService *kafka.Service, dbService *postgres.DbService) web.Routes {
	return web.Routes{
		web.Route{
			Name:    "GetEvents",
			Method:  "GET",
			Pattern: "/events/{domain}",
			HandlerFunc: func(writer http.ResponseWriter, request *http.Request) {
				// TODO get events from database
				// make sure to include
				// dbService.GetAllEventsForDomain()
			},
		},
		web.Route{
			Name:    "PutEvents",
			Method:  "PUT",
			Pattern: "/events/{domain}",
			HandlerFunc: func(writer http.ResponseWriter, request *http.Request) {
				//kafkaService.Publish()
				// TODO put a new event

				//dbService.InsertEvent()
			},
		},
	}
}
