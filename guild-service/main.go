package main

import (
	"fmt"
	"guild-service/application"
	"guild-service/domain"
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

	guildRepository := repositories.NewPostgresGuildRepository(db)
	guildService := application.NewGuildService(guildRepository)

	guild, err := guildService.CreateGuild("Test Guild", "123")
	if err != nil {
		panic(err)
	}

	fmt.Println(guild)

}
