package handlers

import (
	"guild-service/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListGuildsHandler struct {
	guildService domain.GuildService
}

func NewListGuildsHandler(guildService domain.GuildService) *ListGuildsHandler {
	return &ListGuildsHandler{
		guildService: guildService,
	}
}

func (h *ListGuildsHandler) Handle(c *gin.Context) {
	guilds, err := h.guildService.GetAllGuilds()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, guilds)
}
