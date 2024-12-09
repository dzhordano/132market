package config

import (
	"os"
	"time"
)

const (
	jwtSigningKey   = "JWT_SIGNING_KEY"
	accessTokenTTL  = "ACCESS_TOKEN_TTL"
	refreshTokenTTL = "REFRESH_TOKEN_TTL"
)

type JWTConfig interface {
	SigningKey() string
	ATTL() time.Duration
	RTTL() time.Duration
}

type jwtConfig struct {
	signingKey string
	attl       time.Duration
	rttl       time.Duration
}

func (c *jwtConfig) SigningKey() string {
	return c.signingKey
}

func (c *jwtConfig) ATTL() time.Duration {
	return c.attl
}

func (c *jwtConfig) RTTL() time.Duration {
	return c.rttl
}

func MustNewJwtConfig() JWTConfig {
	signingKey := os.Getenv(jwtSigningKey)
	if signingKey == "" {
		panic("JWT_SIGNING_KEY is not set")
	}

	attl := os.Getenv(accessTokenTTL)
	rttl := os.Getenv(refreshTokenTTL)

	convATTL, err := time.ParseDuration(attl)
	if err != nil {
		panic("ACCESS_TOKEN_TTL not set or invalid type")
	}

	convRTTL, err := time.ParseDuration(rttl)
	if err != nil {
		panic("REFRESH_TOKEN_TTL not set or invalid type")
	}

	return &jwtConfig{
		signingKey: signingKey,
		attl:       convATTL,
		rttl:       convRTTL,
	}
}
