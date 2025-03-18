package handlers

import (
	"net/http"

	"api-gateway/domain"

	"github.com/gin-gonic/gin"
)

type VerifyEnrollmentHandler struct {
	enrollmentService domain.EnrollmentService
}

func NewVerifyEnrollmentHandler(enrollmentService domain.EnrollmentService) *VerifyEnrollmentHandler {
	return &VerifyEnrollmentHandler{enrollmentService: enrollmentService}
}

func (h *VerifyEnrollmentHandler) Handle(c *gin.Context) {
	enrollmentID := c.Param("enrollment_id")

	enrollment, err := h.enrollmentService.FindByID(enrollmentID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Enrollment not found"})
		return
	}

	c.JSON(http.StatusOK, enrollment)
}
