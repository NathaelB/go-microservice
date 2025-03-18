package domain

type CourseNotFoundEvent struct {
	ID string `json:"id"`
}

type EnrollmentFailuresEvent struct {
	ID string `json:"id"`
}
