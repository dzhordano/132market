package repository

import (
	"context"
	"time"

	"github.com/dzhordano/132market/services/users/internal/domain/entities"

	"github.com/google/uuid"
)

type UserRepository interface {
	Save(ctx context.Context, user *entities.User) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) (*entities.User, error) // TODO Нужна ли отдельная структура для обновления?
	UpdateLastSeen(ctx context.Context, id uuid.UUID, lastSeen time.Time) error
	UpdateState(ctx context.Context, id uuid.UUID, state string) error
	Delete(ctx context.Context, id uuid.UUID, deletedAt time.Time) error
	FindById(ctx context.Context, id uuid.UUID) (*entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	FindAll(ctx context.Context, offset, limit uint64, filters map[string]string) ([]*entities.User, error)
}
