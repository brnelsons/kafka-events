package postgres

import (
	"kafka-events/models"
)

// language=PostgreSQL
const (
	InsertEvent = `INSERT INTO events (type_uuid, name, description, start_date_time, end_date_time) VALUES ($1, $2, $3, $4, $5)`
	UpdateEvent = `UPDATE events SET type_uuid=$2, name=$3, description=$4, start_date_time=$5, end_date_time=$6 where uuid = $1`
	DeleteEvent = `DELETE FROM events where uuid = $1`
)

func (service *DbService) InsertEvent(event models.Event) error {
	return service.execute(
		InsertEvent,
		event.Type.Uuid,
		event.Name,
		event.Description,
		event.StartDateTime.Unix(),
		event.EndDateTime.Unix(),
	)
}

func (service *DbService) UpdateEvent(event models.Event) error {
	return service.execute(
		UpdateEvent,
		event.Uuid,
		event.Type.Uuid,
		event.Name,
		event.Description,
		event.StartDateTime.Unix(),
		event.EndDateTime.Unix(),
	)
}

func (service *DbService) DeleteEvent(event models.Event) error {
	return service.execute(DeleteEvent, event.Uuid)
}
