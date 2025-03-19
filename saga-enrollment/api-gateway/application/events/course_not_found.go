package events

import (
	"api-gateway/domain"
	"api-gateway/infrastructure"
	"log"
)

func CourseNotFoundConsumer(client *infrastructure.KafkaClient, service domain.EnrollmentService) {
	go func() {

		log.Println("Starting Course Not Found Event consumer...")

		err := infrastructure.ConsumeMessages[domain.CourseNotFoundEvent](
			client,
			"enrollment-course-not-found",
			"api-gateway",
			func(event domain.CourseNotFoundEvent) error {
				log.Printf("Received course not found event: ID=%s", event.ID)

				err := service.FailedNotFound(event.ID)
				if err != nil {
					log.Printf("Error updating failure reason: %v", err)
					return err
				}

				return nil
			},
		)

		if err != nil {
			log.Printf("Error in Course Not Found Event consumer: %v", err)
		}
	}()

}
