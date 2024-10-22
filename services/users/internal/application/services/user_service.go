package services

import (
	"context"
	"dzhordano/132market/services/users/internal/application/command"
	"dzhordano/132market/services/users/internal/application/interfaces"
	"dzhordano/132market/services/users/internal/domain/entities"
	"dzhordano/132market/services/users/internal/domain/repository"

	"github.com/google/uuid"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) interfaces.UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, userCommand command.CreateUserCommand) (entities.User, error) {
	panic("not implemented")
}

func (s *UserService) UpdateUser(ctx context.Context, user entities.User) (entities.User, error) {
	panic("not implemented")
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	panic("not implemented")
}

func (s *UserService) FindUserById(ctx context.Context, id uuid.UUID) (entities.User, error) {
	panic("not implemented")
}

func (s *UserService) FindAllUsers(ctx context.Context) ([]entities.User, error) {
	panic("not implemented")
}
