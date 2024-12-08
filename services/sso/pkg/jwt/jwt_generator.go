package jwt

import (
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

func (j *JwtGenerator) GenerateToken(user_id string, roles []string) (string, error) {
	claims := jwt.MapClaims{
		"id":    user_id,
		"roles": roles,
		"iat":   time.Now().Unix(),
		"exp":   j.AccessTokenTTL,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return jwtToken.SignedString([]byte(j.SigningKey))
}
