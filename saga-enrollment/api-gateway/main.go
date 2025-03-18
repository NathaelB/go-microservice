package main

import (
	"fmt"
	"log"

	"api-gateway/application"
	"api-gateway/domain"
	"api-gateway/infrastructure"
	"api-gateway/infrastructure/repositories"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/api_gateway"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&domain.Enrollment{})
	fmt.Println("Database migrated")

	kafka := infrastructure.NewKafkaClient([]string{"localhost:19092"})

	enrollmentRepo := repositories.NewPostgresEnrollmentRepository(db)
	enrollmentService := application.NewEnrollmentService(enrollmentRepo, kafka)

	httpServer := application.NewHTTPServer(enrollmentService)

	log.Fatal(httpServer.Start(":3333"))
}
