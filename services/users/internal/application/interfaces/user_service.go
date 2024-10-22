package interfaces

import (
	"context"
	"dzhordano/132market/services/users/internal/application/command"
	"dzhordano/132market/services/users/internal/domain/entities"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, userCommand command.CreateUserCommand) (entities.User, error)
	UpdateUser(ctx context.Context, user entities.User) (entities.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	FindUserById(ctx context.Context, id uuid.UUID) (entities.User, error)
	FindAllUsers(ctx context.Context) ([]entities.User, error)
}
