package repositories

import (
	"role-service/domain"

	"gorm.io/gorm"
)

type PostgresRoleRepository struct {
	db *gorm.DB
}

func NewPostgresRoleRepository(db *gorm.DB) domain.RoleRepository {
	return &PostgresRoleRepository{db: db}
}

func (r *PostgresRoleRepository) Save(guild *domain.Role) error {
	return r.db.Create(guild).Error
}

func (r *PostgresRoleRepository) FindByID(id string) (*domain.Role, error) {
	var role domain.Role
	err := r.db.Where("id = ?", id).First(&role).Error

	return &role, err
}
