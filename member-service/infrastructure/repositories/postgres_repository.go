package repositories

import (
	"member-service/domain"

	"gorm.io/gorm"
)

type PostgresMemberRepository struct {
	db *gorm.DB
}

func NewPostgresMemberRepository(db *gorm.DB) domain.MemberRepository {
	return &PostgresMemberRepository{db: db}
}

func (r *PostgresMemberRepository) Save(member *domain.Member) error {
	return r.db.Create(member).Error
}

func (r *PostgresMemberRepository) FindByID(id string) (*domain.Member, error) {
	var member domain.Member
	err := r.db.Where("id = ?", id).First(&member).Error

	return &member, err
}

func (r *PostgresMemberRepository) FindAll() ([]*domain.Member, error) {
	var members []*domain.Member
	err := r.db.Find(&members).
		Error
	return members, err
}
