package application

import (
	"course-service/domain"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	router        *gin.Engine
	courseService domain.CourseService
}

func NewHTTPServer(courseService domain.CourseService) *HTTPServer {
	router := gin.Default()

	server := &HTTPServer{
		router:        router,
		courseService: courseService,
	}

	server.initHandlers()

	server.setupRoutes()

	return server
}

func (s *HTTPServer) initHandlers() {}

func (s *HTTPServer) setupRoutes() {}

func (s *HTTPServer) Start(addr string) error {
	return s.router.Run(addr)
}
