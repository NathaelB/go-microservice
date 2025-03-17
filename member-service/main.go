package main

import (
	"context"
	"fmt"
	"log"
	"member-service/application"
	"member-service/domain"
	"member-service/infrastructure"
	"member-service/infrastructure/repositories"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Créer un contexte avec annulation pour la gestion gracieuse de l'arrêt
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Configuration de la gestion des signaux d'arrêt
	setupSignalHandler(cancel)

	// Connexion à la base de données
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/member_service"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&domain.Member{})
	fmt.Println("Database migrated")

	// Initialisation de Kafka
	kafka := infrastructure.NewKafkaClient([]string{"localhost:19092"})

	// Initialisation des repositories et services
	memberRepository := repositories.NewPostgresMemberRepository(db)
	memberService := application.NewMemberService(memberRepository, kafka)

	// Démarrage du consommateur Kafka pour les événements de création de guilde
	infrastructure.StartGuildCreatedConsumer(ctx, kafka, memberService)

	// Démarrage du serveur HTTP
	httpServer := application.NewHTTPServer(memberService)

	log.Println("Starting HTTP server on :3334")
	if err := httpServer.Start(":3334"); err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}

// setupSignalHandler configure la gestion des signaux pour l'arrêt gracieux
func setupSignalHandler(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Shutdown signal received, gracefully shutting down...")
		cancel()
		// Attendre avant de quitter pour permettre aux goroutines de se terminer proprement
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()
}
