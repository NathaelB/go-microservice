package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

// KafkaClient handles Kafka message production and consumption
type KafkaClient struct {
	brokers []string
}

// NewKafkaClient creates a new Kafka client
func NewKafkaClient(brokers []string) *KafkaClient {
	return &KafkaClient{
		brokers: brokers,
	}
}

// SendMessage sends a serializable message to the specified Kafka topic
func SendMessage[T any](client *KafkaClient, topic string, message T) error {
	// Create a writer for the specified topic
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  client.brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	// Serialize the message to JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to serialize message: %w", err)
	}

	// Send the message
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

// ConsumeMessages consumes messages from the specified Kafka topic and processes them with the provided handler
func ConsumeMessages[T any](client *KafkaClient, topic string, groupID string, handler func(T) error) error {
	// Create a reader for the specified topic
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  client.brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	defer reader.Close()

	// Start consuming messages
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			return fmt.Errorf("error reading message: %w", err)
		}

		// Deserialize the message
		var data T
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			log.Printf("Error deserializing message: %v", err)
			continue
		}

		// Process the message with the provided handler
		if err := handler(data); err != nil {
			log.Printf("Error processing message: %v", err)
		}
	}
}
