package handlers

import (
	"guild-service/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateGuildRequest struct {
	Name    string `json:"name" binding:"required"`
	OwnerID string `json:"owner_id" binding:"required"`
}

// CreateGuildHandler handles the creation of a new guild
type CreateGuildHandler struct {
	guildService domain.GuildService
}

// NewCreateGuildHandler creates a new instance of CreateGuildHandler
func NewCreateGuildHandler(guildService domain.GuildService) *CreateGuildHandler {
	return &CreateGuildHandler{
		guildService: guildService,
	}
}

// Handle processes the guild creation request
func (h *CreateGuildHandler) Handle(c *gin.Context) {
	var req CreateGuildRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	guild, err := h.guildService.CreateGuild(req.Name, req.OwnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, guild)
}
