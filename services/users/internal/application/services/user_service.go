package services

import (
	"context"
	"dzhordano/132market/services/users/internal/application/command"
	"dzhordano/132market/services/users/internal/application/interfaces"
	"dzhordano/132market/services/users/internal/application/mapper"
	"dzhordano/132market/services/users/internal/application/query"
	"dzhordano/132market/services/users/internal/domain/entities"
	"dzhordano/132market/services/users/internal/domain/repository"
	"dzhordano/132market/services/users/pkg/logger"
	"log/slog"

	"github.com/google/uuid"
)

type UserService struct {
	logger logger.Logger
	repo   repository.UserRepository
}

func NewUserService(logger logger.Logger, repo repository.UserRepository) interfaces.UserService {
	return &UserService{
		logger: logger, repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, userCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
	newUser, err := entities.NewUser(userCommand.Name, userCommand.Email, userCommand.Password)
	if err != nil {
		return nil, err
	}

	if err := newUser.Validate(); err != nil {
		return nil, err
	}

	s.logger.Debug("User created and validated:", slog.Any("user", newUser))

	respUser, err := s.repo.Save(ctx, newUser)
	if err != nil {
		return nil, err
	}

	s.logger.Debug("User saved:", slog.Any("user", respUser))

	result := command.CreateUserCommandResult{
		Result: mapper.NewUserResultFromEntity(respUser),
	}

	return &result, nil
}

func (s *UserService) UpdateUser(ctx context.Context, userCommand *command.UpdateUserCommand) (*command.UpdateUserCommandResult, error) {
	panic("not implemented")
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	panic("not implemented")
}

func (s *UserService) FindUserById(ctx context.Context, id uuid.UUID) (*query.UserQueryResult, error) {
	respUser, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	s.logger.Debug("User found:", slog.Any("user", respUser))

	result := query.UserQueryResult{
		Result: mapper.NewUserResultFromEntity(respUser),
	}

	return &result, nil
}

func (s *UserService) FindAllUsers(ctx context.Context, offset, limit int64) (*query.UserQueryListResult, error) {
	panic("not implemented")
}
