package grpc

import (
	"context"

	"github.com/dzhordano/132market/services/sso/pkg/pb/sso_v1"
)

type AuthorizationController struct {
	sso_v1.UnimplementedAuthorizationV1Server
}

func NewAuthorizationController() *AuthorizationController {
	return &AuthorizationController{}
}

func (c *AuthorizationController) GetUserPermissions(ctx context.Context, request *sso_v1.GetUserPermissionsRequest) (*sso_v1.GetUserPermissionsResponse, error) {
	return &sso_v1.GetUserPermissionsResponse{}, nil
}

func (c *AuthorizationController) GetUserRoles(ctx context.Context, request *sso_v1.GetUserRolesRequest) (*sso_v1.GetUserRolesResponse, error) {
	return &sso_v1.GetUserRolesResponse{}, nil
}

func (c *AuthorizationController) AssignRoleToUser(ctx context.Context, request *sso_v1.AssignRoleToUserRequest) (*sso_v1.AssignRoleToUserResponse, error) {
	return &sso_v1.AssignRoleToUserResponse{}, nil
}

func (c *AuthorizationController) RevokeRoleFromUser(ctx context.Context, request *sso_v1.RevokeRoleFromUserRequest) (*sso_v1.RevokeRoleFromUserResponse, error) {
	return &sso_v1.RevokeRoleFromUserResponse{}, nil
}
