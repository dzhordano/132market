package entities

import (
	"net/mail"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("name", "email")

	if err != nil {
		t.Fatalf("No error expected, got: %s", err)
	}

	if user.Name != "name" {
		t.Errorf("Expected name to be 'name', got: %s", user.Name)
	}

	if user.Email != "email" {
		t.Errorf("Expected email to be 'email', got: %s", user.Email)
	}

	if (user.ID == uuid.UUID{}) {
		t.Errorf("Expected ID to be set, got: %s", user.ID)
	}
}

func FuzzValidate(f *testing.F) {
	testcases := [][]string{
		// Валидные случаи
		{"JohnDoe", "johndoe@example.com"},
		{"Anna-Maria", "anna.maria@example.org"},
		{"Short", "short@example.com"},

		// Имя слишком короткое или длинное
		{"A", "a@example.com"},
		{"SuperLongNameThatExceedsFiftyCharactersLimitExactlyFifty", "valid@example.com"},

		// Невалидные email
		{"ValidName", "missing_at_symbol.com"},
		{"ValidName", "missing_domain@.com"},
		{"ValidName", "missing_tld@example."},
		{"ValidName", "extra_spaces @example.com"},

		// Пароли без обязательных символов
		{"ValidName", "valid@example.com"},
		{"ValidName", "valid@example.com"},
		{"ValidName", "valid@example.com"},
		{"ValidName", "valid@example.com"},
		{"ValidName", "valid@example.com"},

		// Пароли с запрещёнными символами
		{"ValidName", "valid@example.com"},
		{"ValidName", "valid@example.com"},
		{"ValidName", "valid@example.com"},

		// Пустые строки
		{"", "", ""},
		{"", "email@example.com"},
		{"ValidName", "", ""},
		{"", "", "ValidP@ss1"},

		// Смешанные случаи
		{"!@#$%^&*", "invalid@example"},
		{"12345", "email@example.com"},
		{"John", "johndoe@example.com"},
		{"ValidName", "email@example.com"},
		{"ValidName", "valid@example.com"},

		// Крайние значения (границы)
		{"Ab", "a@example.com"},
		{"ThisIsExactlyFiftyCharactersLongNameTestingBoundaryCheck!", "valid@example.com"},
	}

	for _, tc := range testcases {
		f.Add(tc[0], tc[1])
	}

	f.Fuzz(func(t *testing.T, name, email string) {
		u, _ := NewUser(name, email)
		errs := u.Validate()
		if len(errs) > 0 {
			if name == "" || len(name) < 3 || len(name) > 50 || strings.ContainsAny(name, symbols+digits) {
				return
			}
			if _, err := mail.ParseAddress(email); err != nil {
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
		}
	})
}
