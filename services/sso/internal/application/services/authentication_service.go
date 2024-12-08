package services

import (
	"context"

	"github.com/dzhordano/132market/services/sso/internal/application/interfaces"
	"github.com/dzhordano/132market/services/sso/internal/domain/entities"
	"github.com/dzhordano/132market/services/sso/pkg/hasher"
	jwtManager "github.com/dzhordano/132market/services/sso/pkg/jwt"
	"github.com/dzhordano/132market/services/sso/pkg/logger"
)

type AuthenticationService struct {
	log      logger.Logger
	usersSvc interfaces.UsersService
	tokens   jwtManager.JwtGenerator
	hasher   hasher.PasswordHasher
}

func NewAuthenticationService(log logger.Logger, usersSvc interfaces.UsersService, tokens jwtManager.JwtGenerator, hasher hasher.PasswordHasher) interfaces.AuthenticationService {
	return &AuthenticationService{
		log:      log,
		usersSvc: usersSvc,
		tokens:   tokens,
		hasher:   hasher,
	}
}

func (s *AuthenticationService) Register(ctx context.Context, email, password string) error {
	return nil
}

func (s *AuthenticationService) Login(ctx context.Context, email, password string) (*entities.Tokens, error) {
	return nil, nil
}

func (s *AuthenticationService) RefreshTokens(ctx context.Context, refreshToken string) (*entities.Tokens, error) {
	return nil, nil
}
