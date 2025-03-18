package repositories

import (
	"api-gateway/domain"

	"gorm.io/gorm"
)

type PostgresEnrollmentRepository struct {
	db *gorm.DB
}

func NewPostgresEnrollmentRepository(db *gorm.DB) domain.EnrollmentRepository {
	return &PostgresEnrollmentRepository{db: db}
}

func (r *PostgresEnrollmentRepository) FindByID(id string) (*domain.Enrollment, error) {
	var enrollment domain.Enrollment
	if err := r.db.Where("id = ?", id).First(&enrollment).Error; err != nil {
		return nil, err
	}
	return &enrollment, nil
}

func (r *PostgresEnrollmentRepository) Create(dto domain.CreateEnrollmentDatabase) (*domain.Enrollment, error) {
	enrollment := domain.Enrollment{
		StudentID:     dto.StudentID,
		CourseID:      dto.CourseID,
		Status:        dto.Status,
		PaymentMethod: dto.PaymentMethod,
	}
	if err := r.db.Create(&enrollment).Error; err != nil {
		return nil, err
	}
	return &enrollment, nil
}
