package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Name         string
	Email        string
	Password     string
	Role         string
	Verified     bool
	AccountState AccountState
	CreatedAt    time.Time
	LastSeen     time.Time
}

type AccountState struct {
	State State
	Since time.Time
}

type State string

const (
	AccountStateActive      State = "active"
	AccountStateDeactivated State = "deactivated"
	AccountStateBanned      State = "banned"
	AccountStateDeleted     State = "deleted"
)

func (u User) Validate() (errors []error) {

	if err := u.ValidateName(); err != nil {
		errors = append(errors, err)
	}

	if err := u.ValidateEmail(); err != nil {
		errors = append(errors, err)
	}

	if err := u.ValidatePassword(); err != nil {
		errors = append(errors, err)
	}

	return
}

func (u User) ValidateName() error {
	if u.Name == "" {
		return errors.New("name can't be empty")
	}

	if len(u.Name) > 16 {
		return errors.New("name can't be longer than 16 characters")
	}

	return nil
}

func (u User) ValidateEmail() error {
	// TODO Email validation regex

	return nil
}

func (u User) ValidatePassword() error {
	if len(u.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	return nil
}
