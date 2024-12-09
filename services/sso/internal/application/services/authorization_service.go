package services

import (
	"context"

	"github.com/dzhordano/132market/services/sso/internal/application/interfaces"
	"github.com/dzhordano/132market/services/sso/internal/domain/repository"
	"github.com/dzhordano/132market/services/sso/pkg/logger"
)

type AuthorizationService struct {
	log        logger.Logger
	usersSvc   interfaces.UsersService
	repository repository.RolesRepository
}

func NewAuthorizationService(log logger.Logger, usersSvc interfaces.UsersService, repository repository.RolesRepository) interfaces.AuthorizationService {
	return &AuthorizationService{log: log, usersSvc: usersSvc, repository: repository}
}

func (a *AuthorizationService) GetUserRoles(ctx context.Context, userId string) ([]string, error) {
	return nil, nil
}

func (a *AuthorizationService) GetUserPermissions(ctx context.Context, userId string) ([]string, error) {
	return nil, nil
}

func (a *AuthorizationService) AssignRoleToUser(ctx context.Context, userId string, role string) error {
	return nil
}

func (a *AuthorizationService) RevokeRoleFromUser(ctx context.Context, userId string, role string) error {
	return nil
}
