package hasher

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	HashSaltSeparator = ":"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
	Verify(hash, password string) error
}

type argon2Hasher struct {
	// time represents the number of
	// passed over the specified memory.
	time uint32
	// cpu memory to be used.
	memory uint32
	// threads for parallelism aspect
	// of the algorithm.
	threads uint8
	// keyLen of the generate hash key.
	keyLen uint32
	// saltLen the length of the salt used.
	saltLen uint32
}

func NewArgon2Hasher(time, saltLen uint32, memory uint32, threads uint8, keyLen uint32) PasswordHasher {
	return &argon2Hasher{
		time:    time,
		saltLen: saltLen,
		memory:  memory,
		threads: threads,
		keyLen:  keyLen,
	}
}

func (h *argon2Hasher) Hash(password string) (string, error) {
	salt, err := randomSecret(h.saltLen)
	if err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}

	hash := argon2.IDKey([]byte(password), []byte(salt), h.time, h.memory, h.threads, h.keyLen)

	return base64.RawStdEncoding.EncodeToString(hash) + HashSaltSeparator + salt, nil
}

func (h *argon2Hasher) Verify(hash, password string) error {
	sepIdx := strings.LastIndex(hash, HashSaltSeparator)
	if sepIdx == -1 {
		return fmt.Errorf("hash is invalid")
	}

	salt := hash[sepIdx+1:]
	hash = hash[:sepIdx]

	passHash := argon2.IDKey([]byte(password), []byte(salt), h.time, h.memory, h.threads, h.keyLen)

	pass := base64.RawStdEncoding.EncodeToString(passHash)

	if pass != hash {
		return fmt.Errorf("invalid password")
	}

	return nil
}

func randomSecret(length uint32) (string, error) {
	salt := make([]byte, length)

	_, err := rand.Read(salt)
	if err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}

	return base64.RawStdEncoding.EncodeToString(salt), nil
}
