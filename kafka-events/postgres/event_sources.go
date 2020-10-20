package postgres

import "kafka-events/models"

// language=PostgreSQL
const (
	InsertEventSource = `INSERT INTO sources (name, description) VALUES ($1, $2)`
	UpdateEventSource = `UPDATE sources SET name=$2, description=$3 WHERE uuid=$1`
	DeleteEventSource = `DELETE FROM sources where uuid = $1`
)

func (service *DbService) InsertSource(eventSource models.EventSource) error {
	return service.execute(
		InsertEventSource,
		eventSource.Name,
		eventSource.Description,
	)
}

func (service *DbService) UpdateSource(eventSource models.EventSource) error {
	return service.execute(
		UpdateEventSource,
		eventSource.Uuid,
		eventSource.Name,
		eventSource.Description,
	)
}

func (service *DbService) DeleteSource(eventSource models.EventSource) error {
	return service.execute(
		DeleteEventSource,
		eventSource.Uuid,
	)
}
