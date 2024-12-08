package services

import (
	"context"

	"github.com/dzhordano/132market/services/sso/internal/application/interfaces"
	tokenValidator "github.com/dzhordano/132market/services/sso/pkg/jwt"
	"github.com/dzhordano/132market/services/sso/pkg/logger"
)

type TokenValidationService struct {
	log       logger.Logger
	validator tokenValidator.TokenValidator
}

func NewTokenValidationService(log logger.Logger, validator tokenValidator.TokenValidator) interfaces.TokenValidationService {
	return &TokenValidationService{log: log, validator: validator}
}

func (t *TokenValidationService) ValidateToken(ctx context.Context, token string) error {
	_, err := t.validator.ValidateToken(token)
	if err != nil {
		return err
	}

	// FIXME добавить какую-то логику (по сути проверять права пользователя). {Надо другой сервис внедрить?}

	t.log.Info("token is valid")

	return nil
}
