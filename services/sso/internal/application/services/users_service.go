package services

import (
	"context"

	"github.com/dzhordano/132market/services/sso/internal/application/interfaces"
	"github.com/dzhordano/132market/services/sso/internal/domain/entities"
	"github.com/dzhordano/132market/services/sso/internal/domain/repository"
	"github.com/dzhordano/132market/services/sso/pkg/logger"
)

// FIXME need to wrap errors from repository and make them grpc status codes

type UsersService struct {
	log        logger.Logger
	repository repository.UsersRepository
}

func NewUsersService(log logger.Logger, repo repository.UsersRepository) interfaces.UsersService {
	return &UsersService{
		log:        log,
		repository: repo,
	}
}

func (u *UsersService) CreateUser(ctx context.Context, email, password string) error {
	user, err := entities.NewUser(email, password)
	if err != nil {
		return err
	}

	err = u.repository.Save(ctx, user)

	u.log.Info("user created", "id", user.ID.String())

	return err
}

func (u *UsersService) FindById(ctx context.Context, id string) (*entities.User, error) {
	user, err := u.repository.FindById(ctx, id)

	u.log.Info("user found", "id", id)

	return user, err
}

func (u *UsersService) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	user, err := u.repository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	u.log.Info("user found", "email", email)

	return user, err
}

func (u *UsersService) Update(ctx context.Context, user *entities.User) error {
	_, err := u.FindById(ctx, user.ID.String())
	if err != nil {
		return err // TODO check if NotFound
	}

	err = u.repository.Update(ctx, user)

	u.log.Info("user updated", "id", user.ID.String())

	return err
}

func (u *UsersService) Delete(ctx context.Context, id string) error {
	err := u.repository.Delete(ctx, id)

	u.log.Info("user deleted", "id", id)

	return err
}

func (u *UsersService) AssignRoleById(ctx context.Context, id, role string) error {
	err := u.repository.AssignRoleById(ctx, id, role)

	u.log.Info("user role assigned", "id", id, "role", role)

	return err
}

func (u *UsersService) RevokeRoleById(ctx context.Context, id, role string) error {
	err := u.repository.RevokeRoleById(ctx, id, role)

	u.log.Info("user role revoked", "id", id, "role", role)

	return err
}
