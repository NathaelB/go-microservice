package infrastructure

import (
	"context"
	"log"
	"role-service/domain"
)

func StartGuildCreatedConsumer(ctx context.Context, client *KafkaClient, service domain.RoleService) {
	go func() {
		log.Println("Starting Guild Created Event consumer...")

		err := ConsumeMessages[domain.GuildCreatedEvent](
			client,
			"create-guild",   // Topic à écouter
			"role-service", // ID du groupe de consommateurs
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