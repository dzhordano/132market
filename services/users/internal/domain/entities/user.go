package entities

import (
	"fmt"
	"net/mail"
	"strings"
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
	CreatedAt    time.Time
	LastSeenAt   time.Time
	IsDeleted    bool
	DeletedAt    time.Time
}

type Role string

const (
	RoleAdmin     Role = "admin"
	RoleUser      Role = "user"
	RoleModerator Role = "moderator"
)

func (r Role) String() string {
	return string(r)
}

func (u *User) AddUserRoles(role ...Role) {
	u.Roles = append(u.Roles, role...)
}

func (u *User) HasRole(role Role) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}

func (u *User) AddRole(role Role) {
	u.Roles = append(u.Roles, role)
}

type Status string

const (
	StatusOnline  Status = "online"
	StatusOffline Status = "offline"
)

func (s Status) String() string {
	return string(s)
}

func (u *User) ChangeStatus(status Status) {
	u.Status = status
	if status == StatusOnline {
		u.LastSeenAt = time.Now()
	}
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

func (u *User) ChangeState(state State) {
	u.State = state
}

func NewUser(name, email, password string) (*User, error) {
	var user *User

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("error while creating user uuid: %w", err)
	}

	user = &User{
		ID:           id,
		Name:         name,
		Email:        email,
		PasswordHash: password,
		Status:       StatusOffline,
		State:        StateActive,
		CreatedAt:    time.Now(),
		LastSeenAt:   time.Now(),
		IsDeleted:    false,
	}

	user.AddRole(RoleUser)

	return user, nil
}

func (u *User) RolesToStrings() []string {
	var roles []string
	for _, r := range u.Roles {
		roles = append(roles, r.String())
	}
	return roles
}

func (u *User) UpdateName(newName string) error {
	u.Name = newName

	return u.Validate()
}

func (u *User) UpdateEmail(newEmail string) error {
	u.Email = newEmail

	return u.Validate()
}

func (u *User) UpdatePassword(newPassword string) error {
	u.PasswordHash = newPassword

	return u.Validate()
}

const notLetters = `!"#$%&'()*+,-./:;<=>?@[\]^_`
const digits = "0123456789"

// FIXME улучшить валидацию
func (u *User) Validate() error {
	if _, err := uuid.Parse(u.ID.String()); err != nil {
		return fmt.Errorf("user id is invalid")
	}

	if u.Name == "" || strings.Contains(u.Name, notLetters) || strings.Contains(u.Name, digits) || len(u.Name) < 2 {
		return fmt.Errorf("invalid name")
	}

	if u.Email == "" {
		return fmt.Errorf("invalid email address")
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return fmt.Errorf("invalid email address")
	}

	if u.PasswordHash == "" || strings.Contains(u.PasswordHash, "_-") {
		return fmt.Errorf("invalid password. symbols '_', '-' and '=' are not allowed")
	}

	if len(u.PasswordHash) < 8 || !strings.Contains(u.PasswordHash, "!@#$%^&*()+") {
		return fmt.Errorf("password must be at least 8 characters long and contain at least one digit and special symbol")
	}

	return nil
}

func (u *User) DeleteUser() error {
	if u.HasRole(RoleAdmin) {
		return fmt.Errorf("cannot delete admin user") // FIXME норм ваще причина? ха-ха
	}

	u.Status = StatusOffline
	u.State = StateDeleted
	u.IsDeleted = true
	u.DeletedAt = time.Now()

	return nil
}
