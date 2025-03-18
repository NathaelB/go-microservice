package handlers

import (
	"api-gateway/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EnrollStudentHandler struct {
	enrollmentService domain.EnrollmentService
}

func NewEnrollStudentHandler(enrollmentService domain.EnrollmentService) *EnrollStudentHandler {
	return &EnrollStudentHandler{enrollmentService: enrollmentService}
}

func (h *EnrollStudentHandler) Handle(c *gin.Context) {
	var req domain.CreateEnrollmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	enrollment, err := h.enrollmentService.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, enrollment)
}
