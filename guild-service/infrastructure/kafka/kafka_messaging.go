package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"guild-service/application"
)

type KafkaMessaging struct {
	producer *kafka.Producer
	config   *kafka.ConfigMap
}

func NewKafkaMessaging(bootstrapServers string) (*KafkaMessaging, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"client.id":         "guild-service",
		"acks":              "all",
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka producer: %w", err)
	}

	return &KafkaMessaging{
		producer: producer,
		config:   config,
	}, nil
}

func (k *KafkaMessaging) PublishMessage(ctx context.Context, topic string, message interface{}) error {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value:     messageBytes,
		Timestamp: time.Now(),
	}

	deliveryChan := make(chan kafka.Event)
	err = k.producer.Produce(kafkaMessage, deliveryChan)
	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}

	select {
	case e := <-deliveryChan:
		m := e.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			return fmt.Errorf("delivery failed: %w", m.TopicPartition.Error)
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (k *KafkaMessaging) Subscribe(ctx context.Context, topic string, groupID string, options application.SubscriptionOptions, handler interface{}) error {
	bootstrapServers, err := k.config.Get("bootstrap.servers", "")
	if err != nil {
		return fmt.Errorf("failed to get bootstrap servers: %w", err)
	}

	consumerConfig := &kafka.ConfigMap{
		"bootstrap.servers":  bootstrapServers,
		"group.id":           groupID,
		"auto.offset.reset":  getOffsetResetValue(options.Offset),
		"enable.auto.commit": true,
	}

	// Create consumer
	consumer, err := kafka.NewConsumer(consumerConfig)
	if err != nil {
		return fmt.Errorf("failed to create consumer: %w", err)
	}
	defer consumer.Close()

	// Subscribe to topic
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return fmt.Errorf("failed to subscribe to topic: %w", err)
	}

	// Start consuming messages
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := consumer.ReadMessage(100 * time.Millisecond)
				if err != nil {
					// Timeout or error
					if err.(kafka.Error).Code() != kafka.ErrTimedOut {
						fmt.Printf("Consumer error: %v\n", err)
					}
					continue
				}

				// Process message based on handler type
				// This is a simplified example - you would need to implement proper type handling
				if messageHandler, ok := handler.(func([]byte) error); ok {
					if err := messageHandler(msg.Value); err != nil {
						fmt.Printf("Handler error: %v\n", err)
					}
				} else {
					fmt.Printf("Unsupported handler type\n")
				}
			}
		}
	}()

	// Wait for context cancellation
	<-ctx.Done()
	return nil
}

// Helper function to convert our Offset type to Kafka's auto.offset.reset value
func getOffsetResetValue(offset application.Offset) string {
	switch offset {
	case application.Beginning:
		return "earliest"
	case application.Latest:
		return "latest"
	default:
		return "latest" // Default to latest
	}
}

// Close closes the Kafka producer
func (k *KafkaMessaging) Close() {
	k.producer.Close()
}
