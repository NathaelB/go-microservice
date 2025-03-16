package main

import (
	"fmt"
	"net/http"
	"role-service/application"
	"role-service/domain"
	"role-service/infrastructure/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func main() {

	router := gin.Default()
	router.Use(func(ctx *gin.Context) {

		db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/role_service"), &gorm.Config{})

		if err != nil {
			panic("failed to connect database")
		}

		db.AutoMigrate(&domain.Role{})

		fmt.Println("Database migrated")
		roleRepository := repositories.NewPostgresRoleRepository(db)
		roleService := application.NewRoleService(roleRepository)
		ctx.Set("roleService", roleService)
		ctx.Next()
	})

	router.POST("/roles", func(c *gin.Context) {

		roleService, _ := c.Get("roleService")
		myService := roleService.(domain.RoleService)
		var newRole *domain.Role

		if err := c.BindJSON(&newRole); err != nil {
			return
		}

		newRole, err := myService.CreateRole(newRole.Name, newRole.GuildID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusCreated, newRole)
	},
	)
	router.GET("/roles/:id", func(c *gin.Context) {
		roleService, _ := c.Get("roleService")
		myService := roleService.(domain.RoleService)
		role, err := myService.GetRoleByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, role)
	})

	router.Run("localhost:8082")

	// curl http://localhost:8082/roles \
	// --include \
	// --header "Content-Type: application/json" \
	// --request "POST" \
	// --data '{"name": "test","guild_id": "1"}'

}
