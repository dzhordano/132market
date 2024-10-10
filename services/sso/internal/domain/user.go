package domain

import (
	"time"

	"github.com/google/uuid"
)

// TODO Переработать чтобы было только то, что нужно для этого сервиса.

type User struct {
	ID           uuid.UUID
	Name         string
	Email        string
	Password     string
	Role         string
	Verified     bool
	Status       string
	AccountState AccountState
	CreatedAt    time.Time
	LastSeen     time.Time
	Session      SessionInfo
}

// Those should be hashed in db i guess
type SessionInfo struct {
	AccessToken  string
	ATExpiresAt  time.Time
	RefreshToken string
	RTExpiresAt  time.Time
}

type AccountState struct {
	State State
	At    time.Time
}

type State string

const (
	AccountStateActive      State = "active"
	AccountStateDeactivated State = "deactivated"
	AccountStateBanned      State = "banned"
	AccountStateDeleted     State = "deleted"
)
