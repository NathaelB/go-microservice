package main

import (
	"fmt"
	"log"
	"role-service/application"
	"role-service/domain"
	"role-service/infrastructure"
	"role-service/infrastructure/repositories"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/role_service"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.Role{})
	fmt.Println("Database migrated")

	kafka := infrastructure.NewKafkaClient([]string{"localhost:19092"})
	roleRepository := repositories.NewPostgresRoleRepository(db)
	roleService := application.NewRoleService(roleRepository, kafka)

	httpServer := application.NewHTTPServer(roleService)

	log.Fatal(httpServer.Start(":3333"))
}
