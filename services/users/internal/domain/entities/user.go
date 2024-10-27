package entities

import (
	"fmt"
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
	DeletedAt    *time.Time
	Tokens       Tokens // TODO мб нужно убрать, когда продумаю JWT с рефрешем над решить оно тута нужновое?
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

type Tokens struct {
	AccessToken  string
	ATExpiresAt  time.Time
	RefreshToken string
	RTExpiresAt  time.Time
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

// FIXME улучшить валидацию
func (u *User) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("user name is empty")
	}
	if u.Email == "" {
		return fmt.Errorf("user email is empty")
	}
	if u.PasswordHash == "" {
		return fmt.Errorf("user password is empty")
	}

	return nil
}
