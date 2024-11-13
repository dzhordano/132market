package interfaces

import (
	"context"
	"dzhordano/132market/services/users/internal/application/command"
	"dzhordano/132market/services/users/internal/application/query"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, userCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error)
	UpdateUser(ctx context.Context, userCommand *command.UpdateUserCommand) (*command.UpdateUserCommandResult, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	FindUserById(ctx context.Context, id uuid.UUID) (*query.UserQueryResult, error)
	FindUserByCredentials(ctx context.Context, email, password string) (*query.UserQueryResult, error)
	FindAllUsers(ctx context.Context, offset, limit uint64) (*query.UserQueryListResult, error)
}
