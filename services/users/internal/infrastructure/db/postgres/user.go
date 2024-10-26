package postgres

import (
	"context"
	"dzhordano/132market/services/users/internal/domain/entities"
	"dzhordano/132market/services/users/internal/domain/repository"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) repository.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindById(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	selectBuilder := sq.Select("*").
		From("users").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(ctx, query, args...)

	var user entities.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Roles, &user.Status, &user.State, &user.CreatedAt, &user.LastSeenAt); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]*entities.User, error) {
	return nil, nil
}

func (r *UserRepository) Save(ctx context.Context, user *entities.User) (*entities.User, error) {
	insertBuilder := sq.Insert("users").
		Columns("id", "name", "email", "password_hash", "roles", "status", "state", "created_at", "last_seen_at").
		Values(user.ID, user.Name, user.Email, user.PasswordHash, user.RolesToStrings(), user.Status, user.State, user.CreatedAt, user.LastSeenAt).
		PlaceholderFormat(sq.Dollar)

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return r.FindById(ctx, user.ID)
}

func (r *UserRepository) Update(ctx context.Context, user *entities.User) (*entities.User, error) {
	return nil, nil
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
