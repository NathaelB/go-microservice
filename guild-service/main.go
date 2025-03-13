package main

import (
	"context"
	"encoding/json"
	"fmt"
	"guild-service/application"
	"guild-service/domain"
	"guild-service/infrastructure/kafka"
	"guild-service/infrastructure/repositories"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/guild_service"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.Guild{})

	fmt.Println("Database migrated")

	kafkaMessaging, err := kafka.NewKafkaMessaging("localhost:9092")

	if err != nil {
		panic("failed to create Kafka messaging: " + err.Error())
	}

	// Define a message type that matches the expected JSON structure
	type TestMessage struct {
		Message string `json:"message"`
	}

	// Create a handler function for the messages
	handler := func(message []byte) error {
		var testMsg TestMessage
		if err := json.Unmarshal(message, &testMsg); err != nil {
			fmt.Printf("Error unmarshaling message: %v\n", err)
			return err
		}
		fmt.Printf("Received message: %s\n", testMsg.Message)
		return nil
	}

	// Start a goroutine to subscribe to the "test" topic
	go func() {
		ctx := context.Background()
		options := application.SubscriptionOptions{
			Offset: application.Latest,
		}
		if err := kafkaMessaging.Subscribe(ctx, "test", "guild-service-group", options, handler); err != nil {
			fmt.Printf("Error subscribing to topic: %v\n", err)
		}
	}()

	fmt.Println("Subscribed to 'test' topic")

	guildRepository := repositories.NewPostgresGuildRepository(db)
	guildService := application.NewGuildService(guildRepository)

	guild, err := guildService.CreateGuild("Test Guild", "123")
	if err != nil {
		panic(err)
	}

	fmt.Println(guild)

}
