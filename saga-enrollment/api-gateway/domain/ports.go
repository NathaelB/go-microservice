package domain

type EnrollmentService interface {
	FindByID(id string) (*Enrollment, error)
	Create(dto CreateEnrollmentRequest) (*Enrollment, error)
}

type EnrollmentRepository interface {
	FindByID(id string) (*Enrollment, error)
	Create(dto CreateEnrollmentDatabase) (*Enrollment, error)
}
