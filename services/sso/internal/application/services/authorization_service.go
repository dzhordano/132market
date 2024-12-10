package services

import (
	"context"
	"errors"

	"github.com/dzhordano/132market/services/sso/internal/application/interfaces"
	"github.com/dzhordano/132market/services/sso/internal/domain/entities"
	"github.com/dzhordano/132market/services/sso/internal/domain/repository"
	"github.com/dzhordano/132market/services/sso/pkg/jwt"
	"github.com/dzhordano/132market/services/sso/pkg/logger"
)

type AuthorizationService struct {
	log            logger.Logger
	usersSvc       interfaces.UsersService
	tokenValidator jwt.TokenValidator
	repository     repository.RolesRepository
}

func NewAuthorizationService(log logger.Logger, usersSvc interfaces.UsersService, tokenValidator jwt.TokenValidator, repository repository.RolesRepository) interfaces.AuthorizationService {
	return &AuthorizationService{log: log, usersSvc: usersSvc, tokenValidator: tokenValidator, repository: repository}
}

func (a *AuthorizationService) GetUserRoles(ctx context.Context, token string) ([]string, error) {
	claims, err := a.tokenValidator.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	rolesIntf := claims["roles"].([]interface{})

	var roles []string
	for _, r := range rolesIntf {
		roles = append(roles, r.(string))
	}

	return roles, nil
}

func (a *AuthorizationService) GetUserPermissions(ctx context.Context, token string) ([]string, error) {
	claims, err := a.tokenValidator.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	rolesIntf := claims["roles"].([]interface{})

	var roles []string
	for _, r := range rolesIntf {
		roles = append(roles, r.(string))
	}

	permissions, err := a.repository.RolesPermissions(ctx, roles)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (a *AuthorizationService) AssignRoleToUser(ctx context.Context, userId string, role string) error {
	if !entities.ValidRole(role) {
		return errors.New("invalid role")
	}

	return a.usersSvc.AssignRoleById(ctx, userId, role)
}

func (a *AuthorizationService) RevokeRoleFromUser(ctx context.Context, userId string, role string) error {
	return a.usersSvc.RevokeRoleById(ctx, userId, role)
}
