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
	tokens   jwtManager.TokenManager
	hasher   hasher.PasswordHasher
}

func NewAuthenticationService(log logger.Logger, usersSvc interfaces.UsersService, tokens jwtManager.TokenManager, hasher hasher.PasswordHasher) interfaces.AuthenticationService {
	return &AuthenticationService{
		log:      log,
		usersSvc: usersSvc,
		tokens:   tokens,
		hasher:   hasher,
	}
}

func (s *AuthenticationService) Register(ctx context.Context, email, password string) error {
	hashedPass, err := s.hasher.Hash(password)
	if err != nil {
		return err
	}

	if err := s.usersSvc.CreateUser(ctx, email, hashedPass); err != nil {
		return err
	}

	// FIXME save in users microservice

	return nil
}

func (s *AuthenticationService) Login(ctx context.Context, email, password string) (*entities.Tokens, error) {
	user, err := s.usersSvc.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// TODO check password hash in users microservice

	accessToken, err := s.tokens.GenerateAccessToken(user.ID.String(), user.Roles)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tokens.GenerateRefreshToken(user.ID.String(), user.Roles)
	if err != nil {
		return nil, err
	}

	return &entities.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthenticationService) RefreshTokens(ctx context.Context, refreshToken string) (*entities.Tokens, error) {
	claims, err := s.tokens.ValidateToken(refreshToken)
	if err != nil {
		return nil, err
	}

	user, err := s.usersSvc.FindById(ctx, claims["sub"].(string))
	if err != nil {
		return nil, err
	}

	accessToken, err := s.tokens.GenerateAccessToken(user.ID.String(), user.Roles)
	if err != nil {
		return nil, err
	}

	refreshToken, err = s.tokens.GenerateRefreshToken(user.ID.String(), user.Roles)
	if err != nil {
		return nil, err
	}

	return &entities.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
