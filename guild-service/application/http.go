package application

import (
	"guild-service/application/handlers"
	"guild-service/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPServer représente le serveur HTTP de l'application
type HTTPServer struct {
	router       *gin.Engine
	guildService domain.GuildService
	handlers     struct {
		createGuild *handlers.CreateGuildHandler
		listGuilds  *handlers.ListGuildsHandler
	}
}

// NewHTTPServer crée une nouvelle instance du serveur HTTP
func NewHTTPServer(guildService domain.GuildService) *HTTPServer {
	router := gin.Default()

	server := &HTTPServer{
		router:       router,
		guildService: guildService,
	}

	// Initialisation des handlers
	server.initHandlers()

	// Configuration des routes
	server.setupRoutes()

	return server
}

// initHandlers initialise tous les handlers
func (s *HTTPServer) initHandlers() {
	s.handlers.createGuild = handlers.NewCreateGuildHandler(s.guildService)
	s.handlers.listGuilds = handlers.NewListGuildsHandler(s.guildService)
}

// setupRoutes configure les routes de l'API
func (s *HTTPServer) setupRoutes() {
	api := s.router.Group("/v1")
	{
		guilds := api.Group("/guilds")
		{
			guilds.GET("/:id", s.GetGuildHandler)
			guilds.POST("", s.handlers.createGuild.Handle)
			guilds.GET("", s.handlers.listGuilds.Handle)
		}
	}
}

// Start démarre le serveur HTTP
func (s *HTTPServer) Start(addr string) error {
	log.Printf("Starting HTTP server on %s", addr)
	return s.router.Run(addr)
}

// GetGuildHandler récupère une guilde par son ID
func (s *HTTPServer) GetGuildHandler(c *gin.Context) {
	id := c.Param("id")

	guild, err := s.guildService.GetGuildByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guild not found"})
		return
	}

	c.JSON(http.StatusOK, guild)
}
