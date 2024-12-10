package repository

import (
	"context"

	"github.com/dzhordano/132market/services/sso/internal/domain/entities"
)

type UsersRepository interface {
	Save(ctx context.Context, user *entities.User) error
	FindById(ctx context.Context, id string) (*entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id string) error

	AssignRoleById(ctx context.Context, id, role string) error
	RevokeRoleById(ctx context.Context, id, role string) error
}

type RolesRepository interface {
	RolesById(ctx context.Context, id string) ([]string, error)
	RolesPermissions(ctx context.Context, roles []string) ([]string, error)
}
