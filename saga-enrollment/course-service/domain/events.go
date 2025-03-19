package domain

type EnrollmentCreateEvent struct {
	ID            string `json:"id"`
	StudentID     string `json:"student_id"`
	CourseID      string `json:"course_id"`
	PaymentMethod string `json:"payment_method"`
}

type EnrollmentCourseNotFoundEvent struct {
	ID string `json:"id"`
}

type EnrollmentCourseValidatedEvent struct {
	ID string `json:"id"`
}
