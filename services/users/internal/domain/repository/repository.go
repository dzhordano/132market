package repository

import (
	"context"
	"dzhordano/132market/services/users/internal/domain/entities"

	"github.com/google/uuid"
)

type UserRepository interface {
	Save(ctx context.Context, user *entities.User) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) (*entities.User, error) // TODO Нужна ли отдельная структура для обновления?
	Delete(ctx context.Context, id uuid.UUID) error
	FindById(ctx context.Context, id uuid.UUID) (*entities.User, error)
	FindAll(ctx context.Context, offset, limit uint64) ([]*entities.User, error)
}
