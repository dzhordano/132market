package jwt

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtGenerator struct {
	SigningKey      string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func NewJwtGenerator(sk string, aTTL, rTTL time.Duration) TokenGenerator {
	return &JwtGenerator{
		SigningKey:      sk,
		AccessTokenTTL:  aTTL,
		RefreshTokenTTL: rTTL,
	}
}

func (j *JwtGenerator) GenerateAccessToken(user_id string, roles []string) (string, error) {
	claims := jwt.MapClaims{
		"id":    user_id,
		"roles": roles,
		"iat":   time.Now().Unix(),
		"exp":   j.AccessTokenTTL,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return jwtToken.SignedString([]byte(j.SigningKey))
}

func (j *JwtGenerator) GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 32) // 256-битный токен
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
