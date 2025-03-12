package repositories

import (
	"member-service/domain"

	"gorm.io/gorm"
)

type PostgresMemberRepository struct {
	db *gorm.DB
}

func NewPostgresMemebrRepository(db *gorm.DB) domain.MemberRepository {
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
