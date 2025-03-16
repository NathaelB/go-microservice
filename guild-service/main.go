package main

import (
	"fmt"
	"guild-service/application"
	"guild-service/domain"
	"guild-service/infrastructure"
	"guild-service/infrastructure/repositories"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Connexion à la base de données
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/guild_service"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migration automatique
	db.AutoMigrate(&domain.Guild{})
	fmt.Println("Database migrated")

	// Init kafka struct

	kafka := infrastructure.NewKafkaClient([]string{"localhost:19092"})

	// Initialisation des repositories et services
	guildRepository := repositories.NewPostgresGuildRepository(db)
	guildService := application.NewGuildService(guildRepository, kafka)

	// Création du serveur HTTP
	httpServer := application.NewHTTPServer(guildService)

	// Démarrage du serveur HTTP
	log.Fatal(httpServer.Start(":3333"))
}
