package domain

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	GuildID   string   `json:"guild_id"`
	MembersID []string `json:"members"`
}

func NewRole(name string, guildID string) (*Role, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	uid := uuid.New().String()

	return &Role{
		ID:      uid,
		Name:    name,
		GuildID: guildID,
	}, nil
}
