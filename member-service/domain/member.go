package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewMember(name string) (*Member, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	uid := uuid.New().String()

	return &Member{
		ID:        uid,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
