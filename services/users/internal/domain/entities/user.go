package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Name         string
	Email        string
	PasswordHash string
	Roles        []Role
	Status       Status
	State        State
	Tokens       Tokens
	CreateaAt    time.Time
	LastSeenAt   time.Time
}

type Role string

const (
	RoleAdmin     Role = "admin"
	RoleUser      Role = "user"
	RoleModerator Role = "moderator"
)

func (u *User) AddUserRoles(role ...Role) {
	u.Roles = append(u.Roles, role...)
}

type Status string

const (
	StatusOnline  Status = "online"
	StatusOffline Status = "offline"
)

type State string

const (
	StateActive  State = "active"
	StateBlocked State = "blocked"
)

type Tokens struct {
	AccessToken  string
	ATExpiresAt  time.Time
	RefreshToken string
	RTExpiresAt  time.Time
}

func (u *User) HasRole(role Role) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}

func (u *User) ChangeStatus(status Status) {
	u.Status = status
	if status == StatusOnline {
		u.LastSeenAt = time.Now()
	}
}

func (u *User) newUser(name, email, password string) *User {
	var user *User

	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}

	user = &User{
		ID:           id,
		Name:         name,
		Email:        email,
		PasswordHash: password,
		Status:       StatusOnline,
		State:        StateActive,
		Tokens:       Tokens{}, // TODO мб поменять...
		CreateaAt:    time.Now(),
		LastSeenAt:   time.Now(),
	}

	return user
}

func (u *User) NewUserWithRoles(name, email, password string, roles ...string) *User {
	castedRoles := make([]Role, 0, len(roles))
	for _, role := range roles {
		castedRoles = append(castedRoles, Role(role))
	}

	u.AddUserRoles(castedRoles...)

	return u.newUser(name, email, password)
}
