package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtValidator struct {
	SigningKey      string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func NewJwtValidator(sk string, aTTL, rTTL time.Duration) TokenValidator {
	return &JwtValidator{
		SigningKey:      sk,
		AccessTokenTTL:  aTTL,
		RefreshTokenTTL: rTTL,
	}
}

func (j *JwtValidator) ValidateToken(token string) (map[string]interface{}, error) {
	claims, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		claims, ok := t.Claims.(jwt.MapClaims)
		if !ok {
			return nil, jwt.ErrInvalidKey
		}

		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			return nil, jwt.ErrTokenExpired
		}

		if claims["iat"].(float64) > float64(time.Now().Unix()) {
			return nil, jwt.ErrInvalidKey
		}

		return []byte(j.SigningKey), nil
	})
	if err != nil {
		return nil, err
	}

	return claims.Claims.(jwt.MapClaims), nil
}
