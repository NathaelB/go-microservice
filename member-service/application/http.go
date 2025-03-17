package application

import (
	"member-service/application/handlers"
	"member-service/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	router        *gin.Engine
	memberService domain.MemberService
	handlers      struct {
		createMember *handlers.CreateMemberHandler
		listMembers  *handlers.ListMembersHandler
	}
}

func NewHTTPServer(memberService domain.MemberService) *HTTPServer {
	router := gin.Default()

	server := &HTTPServer{
		router:        router,
		memberService: memberService,
	}

	server.initHandlers()

	server.setupRoutes()

	return server
}

func (s *HTTPServer) initHandlers() {
	s.handlers.createMember = handlers.NewCreateMemberHandler(s.memberService)
	s.handlers.listMembers = handlers.NewListMembersHandler(s.memberService)
}

func (s *HTTPServer) setupRoutes() {
	api := s.router.Group("/v1")
	{
		members := api.Group("/members")
		{
			members.POST("", s.handlers.createMember.Handle)
			members.GET("", s.handlers.listMembers.Handle)
		}
	}
}

func (s *HTTPServer) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *HTTPServer) GetMemberHandler(c *gin.Context) {
	id := c.Param("id")
	member, err := s.memberService.GetMemberByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, member)
}
