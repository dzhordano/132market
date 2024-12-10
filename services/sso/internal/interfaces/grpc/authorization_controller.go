package grpc

import (
	"context"

	"github.com/dzhordano/132market/services/sso/internal/application/interfaces"
	"github.com/dzhordano/132market/services/sso/pkg/pb/sso_v1"
)

type AuthorizationController struct {
	authorizationService interfaces.AuthorizationService
	sso_v1.UnimplementedAuthorizationV1Server
}

func NewAuthorizationController(as interfaces.AuthorizationService) *AuthorizationController {
	return &AuthorizationController{
		authorizationService: as,
	}
}

func (c *AuthorizationController) GetUserPermissions(ctx context.Context, request *sso_v1.GetUserPermissionsRequest) (*sso_v1.GetUserPermissionsResponse, error) {
	perms, err := c.authorizationService.GetUserPermissions(ctx, request.GetToken())
	if err != nil {
		return nil, err
	}

	return &sso_v1.GetUserPermissionsResponse{
		Permissions: perms,
	}, nil
}

func (c *AuthorizationController) GetUserRoles(ctx context.Context, request *sso_v1.GetUserRolesRequest) (*sso_v1.GetUserRolesResponse, error) {
	roles, err := c.authorizationService.GetUserRoles(ctx, request.GetToken())
	if err != nil {
		return nil, err
	}

	return &sso_v1.GetUserRolesResponse{
		Roles: roles,
	}, nil
}

func (c *AuthorizationController) AssignRoleToUser(ctx context.Context, request *sso_v1.AssignRoleToUserRequest) (*sso_v1.AssignRoleToUserResponse, error) {
	err := c.authorizationService.AssignRoleToUser(ctx, request.GetUserId(), request.GetRole())
	if err != nil {
		return nil, err
	}

	return &sso_v1.AssignRoleToUserResponse{
		Assigned: true,
	}, nil
}

func (c *AuthorizationController) RevokeRoleFromUser(ctx context.Context, request *sso_v1.RevokeRoleFromUserRequest) (*sso_v1.RevokeRoleFromUserResponse, error) {
	err := c.authorizationService.RevokeRoleFromUser(ctx, request.GetUserId(), request.GetRole())
	if err != nil {
		return nil, err
	}

	return &sso_v1.RevokeRoleFromUserResponse{
		Revoked: true,
	}, nil
}
