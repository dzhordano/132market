package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/dzhordano/132market/services/sso/internal/domain/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RolesRepository struct {
	db *pgxpool.Pool
}

func NewRolesRepository(db *pgxpool.Pool) repository.RolesRepository {
	return &RolesRepository{db: db}
}

func (r *RolesRepository) RolesById(ctx context.Context, id string) ([]string, error) {
	selectQuery := sq.Select("roles").
		From("users").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectQuery.ToSql()
	if err != nil {
		return nil, err
	}

	var roles []string
	if err := r.db.QueryRow(ctx, query, args...).Scan(&roles); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *RolesRepository) RolesPermissions(ctx context.Context, roleS []string) ([]string, error) {
	selectQuery := sq.Select("permissions").
		From("roles").
		Where(sq.Eq{"name": roleS}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectQuery.ToSql()
	if err != nil {
		return nil, err
	}

	var permissions []string
	if err := r.db.QueryRow(ctx, query, args...).Scan(&permissions); err != nil {
		return nil, err
	}

	return permissions, nil
}
