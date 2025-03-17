package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Guild représente une guilde dans le système
type Guild struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	OwnerID   string    `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewGuild crée une nouvelle instance de Guild
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
