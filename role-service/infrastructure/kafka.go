package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	brokers []string
}

func NewKafkaClient(brokers []string) *KafkaClient {
	return &KafkaClient{
		brokers: brokers,
	}
}

func SendMessage[T any](client *KafkaClient, topic string, message T) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  client.brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to serialize message: %w", err)
	}

	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: jsonData,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to send message to topic %s: %w", topic, err)
	}

	return nil
}

func ConsumeMessages[T any](client *KafkaClient, topic string, groupID string, handler func(T) error) error {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  client.brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			return fmt.Errorf("failed to read message from topic %s: %w", topic, err)
		}

		var message T
		err = json.Unmarshal(m.Value, &message)
		if err != nil {
			return fmt.Errorf("failed to deserialize message: %w", err)
		}

		err = handler(message)
		if err != nil {
			return fmt.Errorf("failed to process message: %w", err)
		}
	}
}