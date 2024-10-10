package usersrepo

import (
	"context"
	"dzhordano/132market/services/users/internal/domain"
	"dzhordano/132market/services/users/internal/repository"
	"dzhordano/132market/services/users/pkg/databases/pg"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

// Constants for table names.
const (
	usersTable = "users"
)

type usersRepo struct {
	db pg.DB
}

func New(db pg.DB) repository.UsersRepo {
	return &usersRepo{
		db: db,
	}
}

func (r *usersRepo) Create(ctx context.Context, user domain.User) (uuid.UUID, error) {
	builder := sq.Insert(usersTable).
		Columns(
			"name",
			"email",
			"password",
			"role",
			"verified",
			"account_state",
			"account_state_since",
			"created_at",
			"last_seen").
		Values(
			user.Name,
			user.Email,
			user.Password,
			user.Role,
			user.Verified,
			user.AccountState.State,
			user.AccountState.Since,
			user.CreatedAt,
			user.LastSeen,
		).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	sql, args, err := builder.ToSql()
	if err != nil {
		return uuid.Nil, err
	}

	err = r.db.QueryRow(ctx, sql, args...).Scan(&user.ID)
	if err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}

func (r *usersRepo) Update(ctx context.Context, user domain.User) error {
	return nil
}

func (r *usersRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (r *usersRepo) Get(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return domain.User{}, nil
}

func (r *usersRepo) List(ctx context.Context) ([]domain.User, error) {
	return []domain.User{}, nil
}
