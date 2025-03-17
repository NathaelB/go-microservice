package infrastructure

import (
	"context"
	"log"
	"member-service/domain"
)

// StartGuildCreatedConsumer démarre un consommateur Kafka pour les événements de création de guilde
func StartGuildCreatedConsumer(ctx context.Context, client *KafkaClient, service domain.MemberService) {
	go func() {
		log.Println("Starting Guild Created Event consumer...")

		err := ConsumeMessages[domain.GuildCreatedEvent](
			client,
			"create-guild",   // Topic à écouter
			"member-service", // ID du groupe de consommateurs
			func(event domain.GuildCreatedEvent) error {
				log.Printf("Received guild created event: ID=%s, Name=%s, OwnerID=%s",
					event.ID, event.Name, event.OwnerID)

				// Traiter l'événement via le service
				if err := service.HandleGuildCreated(event); err != nil {
					log.Printf("Error handling guild created event: %v", err)
					return err
				}

				log.Printf("Successfully processed guild created event for guild %s", event.ID)
				return nil
			},
		)

		if err != nil {
			log.Printf("Error in Guild Created Event consumer: %v", err)
		}
	}()
}
