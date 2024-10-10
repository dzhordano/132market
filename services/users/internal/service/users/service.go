package usersservice

import (
	"context"
	"dzhordano/132market/services/users/internal/domain"
	"dzhordano/132market/services/users/internal/repository"
	"dzhordano/132market/services/users/internal/service"
	"github.com/google/uuid"
)

type usersService struct {
	usersRepo repository.UsersRepo
}

func New(usersRepo repository.UsersRepo) service.UsersService {
	return &usersService{
		usersRepo: usersRepo,
	}
}

func (s *usersService) Create(ctx context.Context, user domain.User) (uuid.UUID, error) {
	return s.usersRepo.Create(ctx, user)
}

func (s *usersService) Update(ctx context.Context, user domain.User) error {
	return nil
}

func (s *usersService) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (s *usersService) Get(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return domain.User{}, nil
}

func (s *usersService) List(ctx context.Context) ([]domain.User, error) {
	return []domain.User{}, nil
}
