package service

import (
	"context"
	"dzhordano/132market/services/users/internal/domain"
	"github.com/google/uuid"
)

type UsersService interface {
	Create(ctx context.Context, user domain.User) (uuid.UUID, error)
	Update(ctx context.Context, user domain.User) error // TODO Нужна ли структура UpdateUserInput?
	Delete(ctx context.Context, id uuid.UUID) error
	Get(ctx context.Context, id uuid.UUID) (domain.User, error)
	List(ctx context.Context) ([]domain.User, error)
}
