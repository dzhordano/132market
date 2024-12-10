package entities

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID
	Name       string
	Email      string
	Roles      []Role
	Status     Status
	State      State
	CreatedAt  time.Time
	LastSeenAt time.Time
	DeletedAt  time.Time
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

func (s State) Validate() bool {
	switch s {
	case StateActive, StateBlocked, StateDeleted:
		return true
	default:
		return false
	}
}

func (u *User) ChangeState(state State) {
	u.State = state
}

func NewUser(name, email string) (*User, error) {
	var user *User

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("error while creating user uuid: %w", err)
	}

	user = &User{
		ID:         id,
		Name:       name,
		Email:      email,
		Status:     StatusOffline,
		State:      StateActive,
		CreatedAt:  time.Now(),
		LastSeenAt: time.Now(),
		DeletedAt:  time.Time{},
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

	return u.ValidateName()
}

func (u *User) UpdateEmail(newEmail string) error {
	u.Email = newEmail

	return u.ValidateEmail()
}

const symbols = `$*()#@!%/`
const digits = "0123456789"

func (u *User) Validate() (errs []error) {

	if _, err := uuid.Parse(u.ID.String()); err != nil {
		errs = append(errs, errors.New("invalid uuid"))
	}

	if err := u.ValidateName(); err != nil {
		errs = append(errs, err)
	}

	if err := u.ValidateEmail(); err != nil {
		errs = append(errs, err)
	}

	return
}

func (u *User) ValidateName() error {
	if u.Name == "" || strings.ContainsAny(u.Name, symbols+digits) || len(u.Name) < 2 || len(u.Name) > 50 {
		return errors.New("invalid name")
	}
	return nil
}

func (u *User) ValidateEmail() error {
	if u.Email == "" {
		return errors.New("invalid email")
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return fmt.Errorf("invalid email: %w", err)
	}

	return nil
}

func (u *User) MarkDeleted() {
	u.Status = StatusOffline
	u.State = StateDeleted
	u.DeletedAt = time.Now()
}
