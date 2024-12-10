package interfaces

import "context"

type AuthorizationService interface {
	// FIXME token are redundant due to ctx utilization???
	// (WHEN I MAKE REQUEST INTERCEPTOR WILL VALIDATE AND TOKEN WILL REMAIN IN CTX?)
	GetUserPermissions(ctx context.Context, token string) ([]string, error)
	GetUserRoles(ctx context.Context, token string) ([]string, error)

	AssignRoleToUser(ctx context.Context, userId string, role string) error
	RevokeRoleFromUser(ctx context.Context, userId string, role string) error
}
