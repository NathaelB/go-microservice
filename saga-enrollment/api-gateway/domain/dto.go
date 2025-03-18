package domain

type CreateEnrollmentRequest struct {
	StudentID     string `json:"student_id" binding:"required"`
	CourseID      string `json:"course_id" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
}

type CreateEnrollmentDatabase struct {
	StudentID     string `json:"student_id"`
	Status        string `json:"status"`
	CourseID      string `json:"course_id"`
	PaymentMethod string `json:"payment_method"`
}

type UpdateEnrollmentSchema struct {
	Status        string `json:"status,omitempty"`
	FailureReason string `json:"failure_reason,omitempty"`
}
