package services

import (
	"context"
	"dzhordano/132market/services/users/internal/application/command"
	"dzhordano/132market/services/users/internal/application/errors"
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
	newUser, err := entities.NewUser(userCommand.Name, userCommand.Email, userCommand.PasswordHash)
	if err != nil {
		return nil, err
	}

	if err := newUser.Validate(); err != nil {
		return nil, errors.ToGRPCError(errors.ErrBadRequest)
	}

	s.logger.Debug("User created and validated:", slog.Any("user", newUser))

	respUser, err := s.repo.Save(ctx, newUser)
	if err != nil {
		s.logger.Debug("Returning error: ", err)
		return nil, errors.ToGRPCError(err)
	}

	s.logger.Debug("User saved:", slog.Any("user", respUser))

	result := command.CreateUserCommandResult{
		Result: mapper.NewUserResultFromEntity(respUser),
	}

	return &result, nil
}

func (s *UserService) UpdateUser(ctx context.Context, userCommand *command.UpdateUserCommand) (*command.UpdateUserCommandResult, error) {
	userId, err := uuid.Parse(userCommand.ID)
	if err != nil {
		return nil, errors.ToGRPCError(errors.ErrBadRequest)
	}

	user, err := s.repo.FindById(ctx, userId)
	if err != nil {
		return nil, err
	}

	s.logger.Debug("User found:", slog.Any("user", user))

	user.UpdateName(userCommand.Name)
	user.UpdateEmail(userCommand.Email)
	user.UpdatePassword(userCommand.PasswordHash)

	if err := user.Validate(); err != nil {
		return nil, errors.ToGRPCError(errors.ErrBadRequest)
	}

	s.logger.Debug("User updated and validated:", slog.Any("user", user))

	respUser, err := s.repo.Update(ctx, user)
	if err != nil {
		return nil, errors.ToGRPCError(errors.ErrBadRequest)
	}

	s.logger.Debug("User updated:", slog.Any("user", respUser))

	result := command.UpdateUserCommandResult{
		Result: mapper.NewUserResultFromEntity(respUser),
	}

	return &result, nil
}

// FIXME Soft-Delete метод.
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	userId, err := uuid.Parse(id)
	if err != nil {
		return errors.ToGRPCError(errors.ErrBadRequest)
	}

	user, err := s.repo.FindById(ctx, userId)
	if err != nil {
		return err
	}

	if err := user.DeleteUser(); err != nil {
		return err
	}

	if _, err := s.repo.Update(ctx, user); err != nil {
		return err
	}

	s.logger.Debug("User deleted (soft):", slog.Any("id", id))

	return nil
}

func (s *UserService) FindUserById(ctx context.Context, id string) (*query.UserQueryResult, error) {
	userId, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.ToGRPCError(errors.ErrBadRequest)
	}

	respUser, err := s.repo.FindById(ctx, userId)
	if err != nil {
		return nil, err
	}

	s.logger.Debug("User found:", slog.Any("user", respUser))

	result := query.UserQueryResult{
		Result: mapper.NewUserResultFromEntity(respUser),
	}

	return &result, nil
}

func (s *UserService) FindUserByCredentials(ctx context.Context, email, passwordhash string) (*query.UserQueryResult, error) {
	respUser, err := s.repo.FindByCredentials(ctx, email, passwordhash)
	if err != nil {
		return nil, err
	}

	s.logger.Debug("User found:", slog.Any("user", respUser))

	result := query.UserQueryResult{
		Result: mapper.NewUserResultFromEntity(respUser),
	}

	return &result, nil
}

func (s *UserService) FindAllUsers(ctx context.Context, offset, limit uint64) (*query.UserQueryListResult, error) {
	respUsers, err := s.repo.FindAll(ctx, offset, limit)
	if err != nil {
		return nil, errors.ToGRPCError(err)
	}
	// FIXME Очередное маг. число, изменить скок выводить
	s.logger.Debug("Users found:", slog.Any("users", respUsers[:min(int(limit), len(respUsers), 5)]))

	result := query.UserQueryListResult{
		Result: mapper.NewUserResultListFromEntities(respUsers),
	}

	return &result, nil
}
