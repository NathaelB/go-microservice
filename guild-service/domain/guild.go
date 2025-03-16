package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Guild struct {
	gorm.Model
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	OwnerID   string    `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewGuild(name, ownerID string) (*Guild, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	if ownerID == "" {
		return nil, errors.New("ownerID is required")
	}

	uid := uuid.New().String()

	return &Guild{
		ID:        uid,
		Name:      name,
		OwnerID:   ownerID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
