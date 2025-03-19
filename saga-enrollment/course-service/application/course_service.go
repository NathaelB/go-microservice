package application

import (
	"course-service/domain"
	"course-service/infrastructure"
)

type CourseServiceImpl struct {
	repo  domain.CourseRepository
	kafka *infrastructure.KafkaClient
}

func NewCourseService(repo domain.CourseRepository, kafka *infrastructure.KafkaClient) domain.CourseService {
	return &CourseServiceImpl{
		repo:  repo,
		kafka: kafka,
	}
}

func (s *CourseServiceImpl) Create(dto domain.CreateCourseRequest) (*domain.Course, error) {
	course, err := s.repo.Create(dto)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (s *CourseServiceImpl) FindByID(id string) (*domain.Course, error) {
	return s.repo.FindByID(id)
}
