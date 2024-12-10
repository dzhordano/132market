package interfaces

import (
	"context"

	"github.com/dzhordano/132market/services/users/internal/application/command"
	"github.com/dzhordano/132market/services/users/internal/application/query"
)

type UserService interface {
	CreateUser(ctx context.Context, userCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error)
	UpdateUser(ctx context.Context, userCommand *command.UpdateUserCommand) (*command.UpdateUserCommandResult, error)
	DeleteUser(ctx context.Context, id string) error
	FindUserById(ctx context.Context, id string) (*query.UserQueryResult, error)
	FindUserByEmail(ctx context.Context, email string) (*query.UserQueryResult, error)
	ListUsers(ctx context.Context, offset, limit uint64, filters map[string]string) (*query.UserQueryListResult, error)

	CheckUserExists(ctx context.Context, email string) (bool, error)
	UpdateLastSeen(ctx context.Context, id string) error
	SetUserState(ctx context.Context, id string, state string) error
}
