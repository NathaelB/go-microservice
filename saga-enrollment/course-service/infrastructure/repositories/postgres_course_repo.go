package repositories

import (
	"course-service/domain"

	"gorm.io/gorm"
)

type PostgresCourseRepository struct {
	db *gorm.DB
}

func NewPostgresCourseRepository(db *gorm.DB) domain.CourseRepository {
	return &PostgresCourseRepository{db: db}
}

func (r *PostgresCourseRepository) Create(dto domain.CreateCourseRequest) (*domain.Course, error) {
	course := domain.Course{
		Title: dto.Title,
		Seats: dto.Seats,
	}
	if err := r.db.Create(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *PostgresCourseRepository) FindByID(id string) (*domain.Course, error) {
	var course domain.Course
	if err := r.db.Where("id = ?", id).First(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}
