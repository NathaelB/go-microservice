package events

import (
	"course-service/domain"
	"course-service/infrastructure"
	"log"
)

func EnrollmentCreateConsumer(client *infrastructure.KafkaClient, service domain.CourseService) {
	go func() {

		log.Println("Starting Enrollment Create Event consumer...")

		err := infrastructure.ConsumeMessages[domain.EnrollmentCreateEvent](
			client,
			"enrollment-requests",
			"course-service",
			func(event domain.EnrollmentCreateEvent) error {
				log.Printf("Received enrollment create event: ID=%s", event.ID)

				// Verify if the course exists
				// Try to find the course by ID
				course, err := service.FindByID(event.CourseID)

				// Handle course not found scenarios (error or nil course)
				if err != nil || course == nil {
					log.Printf("Course not found or error: ID=%s, Error=%v", event.CourseID, err)

					// Create the course not found event
					notFoundEvent := domain.EnrollmentCourseNotFoundEvent{
						ID: event.ID,
					}

					// Notify that the course doesn't exist
					if notifyErr := infrastructure.SendMessage(client, "enrollment-course-not-found", notFoundEvent); notifyErr != nil {
						log.Printf("Error publishing course-not-found event: %v", notifyErr)
						return notifyErr
					}

					// We don't return an error if the course is not found
					// We only return errors from Kafka operations
					return nil
				}

				// Course exists, notify validation
				validatedEvent := domain.EnrollmentCourseValidatedEvent{
					ID: event.ID,
				}

				if notifyErr := infrastructure.SendMessage(client, "enrollment-course-validated", validatedEvent); notifyErr != nil {
					log.Printf("Error publishing course-validated event: %v", notifyErr)
					return notifyErr
				}

				return nil
			},
		)

		if err != nil {
			log.Printf("Error in Enrollment Create Event consumer: %v", err)
		}
	}()
}
