package model

import (
	"time"

	"github.com/google/uuid"
)

type UserResult struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"username"`
	Email      string    `json:"email"`
	Roles      []string  `json:"roles"`
	Status     string    `json:"status"`
	State      string    `json:"state"`
	LastSeenAt time.Time `json:"last_seen_at"`
	CreatedAt  time.Time `json:"created_at"`
}
