package handlers

import (
	"net/http"
	"role-service/domain"

	"github.com/gin-gonic/gin"
)

type ListRolesHandler struct {
	roleService domain.RoleService
}

func NewListRolesHandler(roleService domain.RoleService) *ListRolesHandler {
	return &ListRolesHandler{
		roleService: roleService,
	}
}

func (h *ListRolesHandler) Handle(c *gin.Context) {
	roles, err := h.roleService.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

