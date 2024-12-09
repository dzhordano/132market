package postgres

import (
	"context"

	"github.com/dzhordano/132market/services/sso/internal/domain/entities"
	"github.com/dzhordano/132market/services/sso/internal/domain/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RolesRepository struct {
	db *pgxpool.Pool
}

func NewRolesRepository(db *pgxpool.Pool) repository.RolesRepository {
	return &RolesRepository{db: db}
}

func (r *RolesRepository) RolesById(ctx context.Context, id string) (*[]entities.Role, error) {
	panic("not implemented") // TODO: Implement
}

func (r *RolesRepository) RolePermissions(ctx context.Context, role string) ([]string, error) {
	panic("not implemented") // TODO: Implement
}
