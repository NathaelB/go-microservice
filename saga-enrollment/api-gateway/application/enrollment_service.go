package application

import (
	"api-gateway/domain"
	"api-gateway/infrastructure"
)

type EnrollmentServiceImpl struct {
	repo  domain.EnrollmentRepository
	kafka *infrastructure.KafkaClient
}

func NewEnrollmentService(repo domain.EnrollmentRepository, kafka *infrastructure.KafkaClient) domain.EnrollmentService {
	return &EnrollmentServiceImpl{
		repo:  repo,
		kafka: kafka,
	}
}

func (s *EnrollmentServiceImpl) FindByID(id string) (*domain.Enrollment, error) {
	return s.repo.FindByID(id)
}

func (s *EnrollmentServiceImpl) Create(dto domain.CreateEnrollmentRequest) (*domain.Enrollment, error) {
	enrollment, err := s.repo.Create(domain.CreateEnrollmentDatabase{
		StudentID:     dto.StudentID,
		CourseID:      dto.CourseID,
		PaymentMethod: dto.PaymentMethod,
		Status:        "PENDING",
	})

	if err != nil {
		return nil, err
	}

	err = infrastructure.SendMessage(s.kafka, "enrollment-requests", enrollment)

	if err != nil {
		return nil, err
	}

	return enrollment, nil
}

func (s *EnrollmentServiceImpl) FailedNotFound(id string) error {
	return s.repo.Update(id, domain.UpdateEnrollmentSchema{
		Status:        "FAILED",
		FailureReason: "COURSE_NOT_FOUND",
	})
}
