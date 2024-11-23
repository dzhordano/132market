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
		{"JohnDoe", "johndoe@example.com", "Passw0rd!"},
		{"Anna-Maria", "anna.maria@example.org", "Str0ngP@ss#"},
		{"Short", "short@example.com", "!P@ss1"},

		// Имя слишком короткое или длинное
		{"A", "a@example.com", "ValidP@ss1"},
		{"SuperLongNameThatExceedsFiftyCharactersLimitExactlyFifty", "valid@example.com", "ValidP@ss1"},

		// Невалидные email
		{"ValidName", "missing_at_symbol.com", "ValidP@ss1"},
		{"ValidName", "missing_domain@.com", "ValidP@ss1"},
		{"ValidName", "missing_tld@example.", "ValidP@ss1"},
		{"ValidName", "extra_spaces @example.com", "ValidP@ss1"},

		// Пароли без обязательных символов
		{"ValidName", "valid@example.com", "alllowercase"},
		{"ValidName", "valid@example.com", "ALLUPPERCASE"},
		{"ValidName", "valid@example.com", "12345678"},
		{"ValidName", "valid@example.com", "NoSpecial123"},
		{"ValidName", "valid@example.com", "Short!1"},

		// Пароли с запрещёнными символами
		{"ValidName", "valid@example.com", "Invalid_Password1"},
		{"ValidName", "valid@example.com", "Invalid-Password2"},
		{"ValidName", "valid@example.com", "Invalid=Password3"},

		// Пустые строки
		{"", "", ""},
		{"", "email@example.com", "password"},
		{"ValidName", "", ""},
		{"", "", "ValidP@ss1"},

		// Смешанные случаи
		{"!@#$%^&*", "invalid@example", "12345678"},
		{"12345", "email@example.com", "!Invalid"},
		{"John", "johndoe@example.com", "OnlySpecial!@#$"},
		{"ValidName", "email@example.com", "noSpecialCharacters123"},
		{"ValidName", "valid@example.com", "TooShort1!"},

		// Крайние значения (границы)
		{"Ab", "a@example.com", "Short!1"},
		{"ThisIsExactlyFiftyCharactersLongNameTestingBoundaryCheck!", "valid@example.com", "V@lidPass123"},
	}

	for _, tc := range testcases {
		f.Add(tc[0], tc[1], tc[2])
	}

	f.Fuzz(func(t *testing.T, name, email, password string) {
		u, _ := NewUser(name, email, password)
		err := u.Validate()
		if err != nil {
			if name == "" || len(name) < 3 || len(name) > 50 {
				return
			}
			if _, err := mail.ParseAddress(email); err != nil {
				return
			}
			if password == "" || len(password) < 8 || !strings.ContainsAny(password, symbols) || !strings.ContainsAny(password, digits) {
				return
			}
			for _, ch := range password {
				if !strings.ContainsRune(letters, ch) && !strings.ContainsRune(digits, ch) && !strings.ContainsRune(symbols, ch) {
					return
				}
			}
			t.Errorf("presented error: %s", err)
		} else {
			if len(name) < 2 || len(name) > 50 || strings.ContainsAny(name, symbols+digits) {
				t.Errorf("expected name validation to fail for: %s", name)
			}
			if _, err := mail.ParseAddress(email); err != nil {
				t.Errorf("expected email validation to fail for: %s", email)
			}
			if len(password) < 8 || !strings.ContainsAny(password, symbols) || strings.ContainsAny(password, "_-=") {
				t.Errorf("expected password validation to fail for: %s", password)
			}
		}
	})
}
