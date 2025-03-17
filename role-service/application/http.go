package application

import (
	"net/http"
	"role-service/application/handlers"
	"role-service/domain"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	router       *gin.Engine
	roleService domain.RoleService
	handlers     struct {
		createRole *handlers.CreateRoleHandler
		listRoles  *handlers.ListRolesHandler
	}

}

func NewHTTPServer(roleService domain.RoleService) *HTTPServer {
	router := gin.Default()

	server := &HTTPServer{
		router:       router,
		roleService: roleService,
	}

	server.initHandlers()
	server.setupRoutes()

	return server
}

func (s *HTTPServer) initHandlers() {
	s.handlers.createRole = handlers.NewCreateRoleHandler(s.roleService)
	s.handlers.listRoles = handlers.NewListRolesHandler(s.roleService)
}

func (s *HTTPServer) setupRoutes() {
	api := s.router.Group("/v1")
	{
		roles := api.Group("/roles")
		{
			roles.POST("", s.handlers.createRole.Handle)
			roles.GET("", s.handlers.listRoles.Handle)
		}
	}
}

func (s *HTTPServer) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *HTTPServer) GetRoleHandler(c *gin.Context) {
	roleID := c.Param("id")
	role, err := s.roleService.GetRoleByID(roleID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, role)
}