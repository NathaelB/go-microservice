package main

import (
	"fmt"
	"member-service/application"
	"member-service/domain"
	"member-service/infrastructure/repositories"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/member_service"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.Member{})

	fmt.Println("Database migrated")

	memberRepository := repositories.NewPostgresMemebrRepository(db)
	memberService := application.NewMemberService(memberRepository)

	member, err := memberService.CreateMember("Test Member")
	if err != nil {
		panic(err)
	}

	fmt.Println(member)

}
