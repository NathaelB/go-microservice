package domain

type EnrollmentService interface {
	FindByID(id string) (*Enrollment, error)
	Create(dto CreateEnrollmentRequest) (*Enrollment, error)
	FailedNotFound(id string) error
}

type EnrollmentRepository interface {
	FindByID(id string) (*Enrollment, error)
	Create(dto CreateEnrollmentDatabase) (*Enrollment, error)
	Update(id string, schema UpdateEnrollmentSchema) error
}
