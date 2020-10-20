package models

import "time"

type EventType struct {
	Uuid        string      `json:"uuid"`
	Source      EventSource `json:"source"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
}

type EventSource struct {
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Event struct {
	Uuid          string    `json:"uuid"`
	Type          EventType `json:"type"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	StartDateTime time.Time `json:"startDateTime"`
	EndDateTime   time.Time `json:"endDateTime"`
}

type EventAssociation struct {
	Uuid        string      `json:"uuid"`
	Event       Event       `json:"event"`
	Source      EventSource `json:"source"`
	Description string      `json:"description"`
	DateTime    time.Time   `json:"dateTime"`
}
