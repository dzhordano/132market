package entities

import (
	"net/mail"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("name", "email", "password")

	if err != nil {
		t.Fatalf("No error expected, got: %s", err)
	}

	if user.Name != "name" {
		t.Errorf("Expected name to be 'name', got: %s", user.Name)
	}

	if user.Email != "email" {
		t.Errorf("Expected email to be 'email', got: %s", user.Email)
	}

	if user.PasswordHash != "password" {
		t.Errorf("Expected password to be 'password', got: %s", user.PasswordHash)
	}

	if (user.ID == uuid.UUID{}) {
		t.Errorf("Expected ID to be set, got: %s", user.ID)
	}
}

func FuzzValidate(f *testing.F) {
	testcases := [][]string{
		// Валидные случаи
		{"JohnDoe", "johndoe@example.com", "somepasshash"},
		{"Anna-Maria", "anna.maria@example.org", "somepasshash"},
		{"Short", "short@example.com", "somepasshash"},

		// Имя слишком короткое или длинное
		{"A", "a@example.com", "somepasshash"},
		{"SuperLongNameThatExceedsFiftyCharactersLimitExactlyFifty", "valid@example.com", "somepasshash"},

		// Невалидные email
		{"ValidName", "missing_at_symbol.com", "somepasshash"},
		{"ValidName", "missing_domain@.com", "somepasshash"},
		{"ValidName", "missing_tld@example.", "somepasshash"},
		{"ValidName", "extra_spaces @example.com", "somepasshash"},

		// Пароли без обязательных символов
		{"ValidName", "valid@example.com", "somepasshash"},
		{"ValidName", "valid@example.com", "somepasshash"},
		{"ValidName", "valid@example.com", "somepasshash"},
		{"ValidName", "valid@example.com", "somepasshash"},
		{"ValidName", "valid@example.com", "somepasshash"},

		// Пароли с запрещёнными символами
		{"ValidName", "valid@example.com", "somepasshash"},
		{"ValidName", "valid@example.com", "somepasshash"},
		{"ValidName", "valid@example.com", "somepasshash"},

		// Пустые строки
		{"", "", ""},
		{"", "email@example.com", "somepasshash"},
		{"ValidName", "", ""},
		{"", "", "ValidP@ss1"},

		// Смешанные случаи
		{"!@#$%^&*", "invalid@example", "somepasshash"},
		{"12345", "email@example.com", "somepasshash"},
		{"John", "johndoe@example.com", "somepasshash"},
		{"ValidName", "email@example.com", "somepasshash"},
		{"ValidName", "valid@example.com", "somepasshash"},

		// Крайние значения (границы)
		{"Ab", "a@example.com", "somepasshash"},
		{"ThisIsExactlyFiftyCharactersLongNameTestingBoundaryCheck!", "valid@example.com", "somepasshash"},
	}

	for _, tc := range testcases {
		f.Add(tc[0], tc[1], tc[2])
	}

	f.Fuzz(func(t *testing.T, name, email, password string) {
		u, _ := NewUser(name, email, password)
		errs := u.Validate()
		if len(errs) > 0 {
			if name == "" || len(name) < 3 || len(name) > 50 || strings.ContainsAny(name, symbols+digits) {
				return
			}
			if _, err := mail.ParseAddress(email); err != nil {
				return
			}
			if password == "" {
				return
			}
			t.Errorf("presented error: %s", errs) // fixme weird...
		} else {
			if len(name) < 2 || len(name) > 50 || strings.ContainsAny(name, symbols+digits) {
				t.Errorf("expected name validation to fail for: %s", name)
			}
			if _, err := mail.ParseAddress(email); err != nil {
				t.Errorf("expected email validation to fail for: %s", email)
			}
			if password == "" {
				t.Errorf("expected password validation to fail for: %s", password)
			}
		}
	})
}
