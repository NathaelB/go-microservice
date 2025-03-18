package application

import (
	"log"

	"api-gateway/application/handlers"
	"api-gateway/domain"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	router            *gin.Engine
	enrollmentService domain.EnrollmentService

	handlers struct {
		verifyEnrollmentHandler *handlers.VerifyEnrollmentHandler
		enrollStudentHandler    *handlers.EnrollStudentHandler
	}
}

func NewHTTPServer(enrollmentService domain.EnrollmentService) *HTTPServer {
	router := gin.Default()

	server := &HTTPServer{
		router:            router,
		enrollmentService: enrollmentService,
	}

	server.initHandlers()
	server.setupRoutes()

	return server
}

func (s *HTTPServer) initHandlers() {
	s.handlers.verifyEnrollmentHandler = handlers.NewVerifyEnrollmentHandler(s.enrollmentService)
	s.handlers.enrollStudentHandler = handlers.NewEnrollStudentHandler(s.enrollmentService)
}

func (s *HTTPServer) setupRoutes() {
	api := s.router.Group("/api")
	{
		api.GET("/enrollments/:enrollment_id", s.handlers.verifyEnrollmentHandler.Handle)
		api.POST("/enrollments", s.handlers.enrollStudentHandler.Handle)
	}

}

func (s *HTTPServer) Start(addr string) error {
	log.Printf("Starting HTTP server on %s", addr)
	return s.router.Run(addr)
}
