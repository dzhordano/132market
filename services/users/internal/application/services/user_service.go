package services

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/dzhordano/132market/services/users/internal/application/command"
	svcErrors "github.com/dzhordano/132market/services/users/internal/application/errors"
	"github.com/dzhordano/132market/services/users/internal/application/interfaces"
	"github.com/dzhordano/132market/services/users/internal/application/mapper"
	"github.com/dzhordano/132market/services/users/internal/application/query"
	"github.com/dzhordano/132market/services/users/internal/domain/entities"
	"github.com/dzhordano/132market/services/users/internal/domain/repository"
	"github.com/dzhordano/132market/services/users/internal/infrastructure/db/postgres"
	"github.com/dzhordano/132market/services/users/pkg/logger"

	"github.com/google/uuid"
)

const (
	DefaultOffset uint64 = 0
	DefaultLimit  uint64 = 10

	MaxOffset uint64 = 100
	MaxLimit  uint64 = 100
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

	errs := newUser.Validate()
	if len(errs) > 0 {
		return nil, svcErrors.ToGRPCErrors(svcErrors.ErrBadRequest, errs)
	}

	s.logger.Debug("User created and validated:", slog.Any("user", newUser))

	respUser, err := s.repo.Save(ctx, newUser)
	if err != nil {
		s.logger.Debug("Returning error: ", err)
		return nil, svcErrors.ToGRPCError(err)
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
		return nil, svcErrors.ToGRPCError(svcErrors.ErrBadRequest)
	}

	user, err := s.repo.FindById(ctx, userId)
	if err != nil {
		return nil, err
	}

	s.logger.Debug("User found:", slog.Any("user", user))

	user.UpdateName(userCommand.Name)
	user.UpdateEmail(userCommand.Email)
	user.UpdatePassword(userCommand.PasswordHash)

	errs := user.Validate()
	if len(errs) > 0 {
		return nil, svcErrors.ToGRPCErrors(svcErrors.ErrBadRequest, errs)
	}

	s.logger.Debug("User updated and validated:", slog.Any("user", user))

	respUser, err := s.repo.Update(ctx, user)
	if err != nil {
		return nil, svcErrors.ToGRPCError(svcErrors.ErrBadRequest)
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
		return svcErrors.ToGRPCError(svcErrors.ErrBadRequest)
	}

	if err := s.repo.Delete(ctx, userId, time.Now()); err != nil {
		return err
	}

	s.logger.Debug("User deleted (soft):", slog.Any("id", id))

	return nil
}

func (s *UserService) FindUserById(ctx context.Context, id string) (*query.UserQueryResult, error) {
	userId, err := uuid.Parse(id)
	if err != nil {
		return nil, svcErrors.ToGRPCError(svcErrors.ErrBadRequest)
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

func (s *UserService) FindUserByEmail(ctx context.Context, email string) (*query.UserQueryResult, error) {
	respUser, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	s.logger.Debug("User found:", slog.Any("user", respUser))

	result := query.UserQueryResult{
		Result: mapper.NewUserResultFromEntity(respUser),
	}

	return &result, nil
}

func (s *UserService) ListUsers(ctx context.Context, offset, limit uint64, filters map[string]string) (*query.UserQueryListResult, error) {
	if offset > MaxOffset {
		return nil, svcErrors.ToGRPCError(svcErrors.ErrBadRequest)
	}

	if limit > MaxLimit {
		return nil, svcErrors.ToGRPCError(svcErrors.ErrBadRequest)
	}

	if limit == 0 {
		limit = DefaultLimit
	}

	if offset == 0 {
		offset = DefaultOffset
	}

	respUsers, err := s.repo.FindAll(ctx, offset, limit, filters)
	if err != nil {
		return nil, svcErrors.ToGRPCError(err)
	}

	s.logger.Debug("Users found:", slog.Any("users", respUsers[:min(int(limit), len(respUsers))]))

	result := query.UserQueryListResult{
		Result: mapper.NewUserResultListFromEntities(respUsers),
		Count:  uint64(len(respUsers)),
	}

	return &result, nil
}

func (s *UserService) CheckUserExists(ctx context.Context, email string) (bool, error) {
	_, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, postgres.ErrNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (s *UserService) UpdateLastSeen(ctx context.Context, id string) error {
	userId, err := uuid.Parse(id)
	if err != nil {
		return svcErrors.ToGRPCError(svcErrors.ErrBadRequest)
	}

	if err := s.repo.UpdateLastSeen(ctx, userId, time.Now()); err != nil {
		return err
	}

	s.logger.Debug("User last seen updated:", slog.Any("id", id))

	return nil
}

func (s *UserService) SetUserState(ctx context.Context, id string, state string) error {
	userId, err := uuid.Parse(id)
	if err != nil {
		return svcErrors.ToGRPCError(svcErrors.ErrBadRequest)
	}

	domainState := entities.State(state)
	if !domainState.Validate() {
		return svcErrors.ToGRPCError(svcErrors.ErrBadRequest)
	}

	if err := s.repo.UpdateState(ctx, userId, state); err != nil {
		return err
	}

	s.logger.Debug("User status updated:", slog.Any("id", id), slog.Any("state", state))

	return nil
}
