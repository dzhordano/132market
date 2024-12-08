package interfaces

import "context"

type AuthorizationService interface {
	GetUserPermissions(ctx context.Context, userId string) ([]string, error)
	GetUserRoles(ctx context.Context, userId string) ([]string, error)

	AssignRoleToUser(ctx context.Context, userId string, role string) error
	RevokeRoleFromUser(ctx context.Context, userId string, role string) error
}
