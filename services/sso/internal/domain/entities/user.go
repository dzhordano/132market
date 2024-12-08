package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Email    string
	Password string
	Roles    []Role
	State    State
}

type Role struct {
	Name        string
	Permissions []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

const (
	RoleAdmin string = "admin"
	RoleModer string = "moderator"
	RoleUser  string = "user"
)

func (r User) HasRole(role string) bool {
	for _, v := range r.Roles {
		if v.Name == role {
			return true
		}
	}
	return false
}

type State string

const (
	StateActive  State = "active"
	StateBlocked State = "blocked"
	StateDeleted State = "deleted"
)

func (s State) String() string {
	return string(s)
}

func NewUser(email, password string) (*User, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("error generating uuid for user: %v", err)
	}

	return &User{
		ID:       id,
		Email:    email,
		Password: password,
		State:    StateActive,
	}, nil
}
