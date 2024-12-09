package postgres

import (
	"context"

	"github.com/dzhordano/132market/services/sso/internal/domain/entities"
	"github.com/dzhordano/132market/services/sso/internal/domain/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository struct {
	db *pgxpool.Pool
}

func NewUsersRepository(db *pgxpool.Pool) repository.UsersRepository {
	return &UsersRepository{db: db}
}

func (r *UsersRepository) Save(ctx context.Context, user *entities.User) error {
	panic("not implemented") // TODO: Implement
}

func (r *UsersRepository) FindById(ctx context.Context, id string) (*entities.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r *UsersRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r *UsersRepository) Update(ctx context.Context, user *entities.User) error {
	panic("not implemented") // TODO: Implement
}

func (r *UsersRepository) Delete(ctx context.Context, id string) error {
	panic("not implemented") // TODO: Implement
}

func (r *UsersRepository) AssignRoleById(ctx context.Context, id string, role string) error {
	panic("not implemented") // TODO: Implement
}

func (r *UsersRepository) RevokeRoleById(ctx context.Context, id string, role string) error {
	panic("not implemented") // TODO: Implement
}
