package events

import (
	"course-service/domain"
	"course-service/infrastructure"
	"errors"
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
				course, err := service.FindByID(event.CourseID)
				if err != nil {
					log.Printf("Error finding course: %v", err)

					notifyErr := infrastructure.SendMessage(client, "enrollment-course-not-found", domain.EnrollmentCourseNotFoundEvent{
						ID: event.ID,
					})

					if notifyErr != nil {
						log.Printf("Error publishing course-not-found event: %v", notifyErr)
					}

					return err
				}

				if course == nil {
					log.Printf("Course not found: ID=%s", event.CourseID)

					// Notify that the course doesn't exist
					notifyErr := infrastructure.SendMessage(client, "enrollment-course-not-found", domain.EnrollmentCourseNotFoundEvent{
						ID: event.ID,
					})

					if notifyErr != nil {
						log.Printf("Error publishing course-not-found event: %v", notifyErr)
					}

					return errors.New("course not found")
				}

				// Notify that the course exists
				notifyErr := infrastructure.SendMessage(client, "enrollment-course-validated", domain.EnrollmentCourseValidatedEvent{
					ID: event.ID,
				})

				if notifyErr != nil {
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
