package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type Service struct {
	writer *kafka.Writer
}

func NewKafkaService(topic string, brokerAddresses ...string) *Service {
	w := &kafka.Writer{
		Addr:         kafka.TCP(brokerAddresses...),
		Topic:        topic,
		RequiredAcks: kafka.RequireAll,
		Async:        false,
	}

	return &Service{
		writer: w,
	}
}

func (service Service) PublishWithPartition(context context.Context, partitionKey string, message string) error {
	return service.writer.WriteMessages(context,
		kafka.Message{
			Key:   []byte(partitionKey),
			Value: []byte(message),
		})
}

func (service Service) Publish(context context.Context, message string) error {
	return service.writer.WriteMessages(context,
		kafka.Message{
			Value: []byte(message),
		})
}
