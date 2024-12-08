package config

import (
	"os"
	"strconv"
)

const (
	saltLen    = "ARGON_SALT_LENGTH"
	keyLen     = "ARGON_KEY_LENGTH"
	timeCost   = "ARGON_TIME_COST"
	memoryCost = "ARGON_MEMORY_COST"
	threads    = "ARGON_THREADS"
)

type Argon2Config interface {
	Time() uint32
	SaltLen() uint32
	Memory() uint32
	Threads() uint8
	KeyLen() uint32
}

type argon2Config struct {
	time    uint32
	saltLen uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

func MustNewArgonConfig() Argon2Config {
	return &argon2Config{
		time:    uint32(getEnvInt(timeCost)),
		saltLen: uint32(getEnvInt(saltLen)),
		memory:  uint32(getEnvInt(memoryCost)),
		threads: uint8(getEnvInt(threads)),
		keyLen:  uint32(getEnvInt(keyLen)),
	}
}

func (a *argon2Config) Time() uint32 {
	return a.time
}

func (a *argon2Config) SaltLen() uint32 {
	return a.saltLen
}

func (a *argon2Config) Memory() uint32 {
	return a.memory
}

func (a *argon2Config) Threads() uint8 {
	return a.threads
}

func (a *argon2Config) KeyLen() uint32 {
	return a.keyLen
}

func getEnvInt(key string) int {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is not set")
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return intValue
}
