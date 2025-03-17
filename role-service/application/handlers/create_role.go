package handlers

import (
	"role-service/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateRoleRequest struct {
	Name string `json:"name" binding:"required"`
	GuildID string `json:"guild_id" binding:"required"`

}

type CreateRoleHandler struct {
	roleService domain.RoleService
}

func NewCreateRoleHandler(roleService domain.RoleService) *CreateRoleHandler {
	return &CreateRoleHandler{
		roleService: roleService,
	}
}

func (h *CreateRoleHandler) Handle(c *gin.Context) {
	var req CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := h.roleService.CreateRole(req.Name , req.GuildID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, role)
}