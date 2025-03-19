package main

import (
	"course-service/application"
	"course-service/application/events"
	"course-service/domain"
	"course-service/infrastructure"
	"course-service/infrastructure/repositories"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/course_service"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&domain.Course{}, &domain.Student{}, &domain.CourseStudent{})
	fmt.Println("Database migrated")

	kafka := infrastructure.NewKafkaClient([]string{"localhost:19092"})

	courseRepo := repositories.NewPostgresCourseRepository(db)
	courseService := application.NewCourseService(courseRepo, kafka)

	events.EnrollmentCreateConsumer(kafka, courseService)

	httpServer := application.NewHTTPServer(courseService)
	log.Fatal(httpServer.Start(":3334"))
}
