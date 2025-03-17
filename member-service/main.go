package main

import (
	"fmt"
	"log"
	"member-service/application"
	"member-service/domain"
	"member-service/infrastructure"
	"member-service/infrastructure/repositories"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/member_service"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)

	}

	db.AutoMigrate(&domain.Member{})
	fmt.Println("Database migrated")

	kafka := infrastructure.NewKafkaClient([]string{"localhost:19092"})

	memberRepository := repositories.NewPostgresMemberRepository(db)
	memberService := application.NewMemberService(memberRepository, kafka)

	httpServer := application.NewHTTPServer(memberService)

	log.Fatal(httpServer.Start(":3333"))


}
