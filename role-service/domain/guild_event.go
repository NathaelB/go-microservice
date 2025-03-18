package domain

import (
	"time"
)

// GuildCreatedEvent représente l'événement émis lorsqu'une guilde est créée
type GuildCreatedEvent struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	OwnerID   string    `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
