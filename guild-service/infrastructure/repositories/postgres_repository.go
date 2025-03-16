package repositories

import (
	"guild-service/domain"

	"gorm.io/gorm"
)

type PostgresGuildRepository struct {
	db *gorm.DB
}

func NewPostgresGuildRepository(db *gorm.DB) domain.GuildRepository {
	return &PostgresGuildRepository{db: db}
}

func (r *PostgresGuildRepository) Save(guild *domain.Guild) error {
	return r.db.Create(guild).Error
}

func (r *PostgresGuildRepository) FindByID(id string) (*domain.Guild, error) {
	var guild domain.Guild
	err := r.db.Where("id = ?", id).First(&guild).Error

	return &guild, err
}

func (r *PostgresGuildRepository) FindAll() ([]*domain.Guild, error) {
	var guilds []*domain.Guild
	err := r.db.Find(&guilds).Error
	return guilds, err
}
