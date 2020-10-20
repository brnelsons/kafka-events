package postgres

import (
	"kafka-events/models"
)

// language=PostgreSQL
const (
	InsertEventType = `INSERT INTO event_types (source_uuid, name, description) VALUES ($1, $2, $3)`
	UpdateEventType = `UPDATE event_types SET name = $2, description = $3 where uuid = $1`
	DeleteEventType = `DELETE FROM event_types where uuid = $1`
)

func (service *DbService) InsertEventType(eventType models.EventType) error {
	return service.execute(InsertEventType, eventType.Source.Uuid, eventType.Name, eventType.Description)
}

func (service *DbService) UpdateEventType(eventType models.EventType) error {
	return service.execute(UpdateEventType, eventType.Source.Uuid, eventType.Name, eventType.Description)
}

func (service *DbService) DeleteEventType(eventType models.EventType) error {
	return service.execute(DeleteEventType, eventType.Source.Uuid)
}
