package postgres

import (
	"context"
	"dzhordano/132market/services/users/internal/domain/entities"
	"dzhordano/132market/services/users/internal/domain/repository"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	usersTable = "users"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) repository.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindById(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	const op = "repository.user.FindById"

	selectBuilder := sq.Select("*").
		From(usersTable).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	row := r.db.QueryRow(ctx, query, args...)

	var user entities.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Roles, &user.Status, &user.State, &user.CreatedAt, &user.LastSeenAt, &user.IsDeleted, &user.DeletedAt); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.NoDataFound {
				return nil, fmt.Errorf("%s: %w", op, ErrNotFound)
			}
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}

func (r *UserRepository) FindByCredentials(ctx context.Context, email, password string) (*entities.User, error) {
	const op = "repository.user.FindByCredentials"

	selectBuilder := sq.Select("*").
		From(usersTable).
		Where(sq.And{
			sq.Eq{"email": email},
			sq.Eq{"password_hash": password},
		}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	row := r.db.QueryRow(ctx, query, args...)

	var user entities.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Roles, &user.Status, &user.State, &user.CreatedAt, &user.LastSeenAt, &user.IsDeleted, &user.DeletedAt); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.NoDataFound {
				return nil, fmt.Errorf("%s: %w", op, ErrNotFound) // TODO Сделать здесь константу
			}
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}

func (r *UserRepository) FindAll(ctx context.Context, offset, limit uint64) ([]*entities.User, error) {
	const op = "repository.user.FindAll"

	selectBuilder := sq.Select("*").
		From(usersTable).
		Offset(offset).
		Limit(limit).
		OrderBy("name ASC").
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var users []*entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Roles, &user.Status, &user.State, &user.CreatedAt, &user.LastSeenAt, &user.IsDeleted, &user.DeletedAt); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		users = append(users, &user)
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// FIXME обернуть нормально ошибку
	if len(users) == 0 {
		return nil, fmt.Errorf("%s: %w", op, ErrNotFound) // TODO Сделать константой
	}

	return users, nil
}

func (r *UserRepository) Save(ctx context.Context, user *entities.User) (*entities.User, error) {
	const op = "repository.user.Save"

	insertBuilder := sq.Insert(usersTable).
		Columns("id", "name", "email", "password_hash", "roles", "status", "state", "created_at", "last_seen_at", "is_deleted", "deleted_at").
		Values(user.ID, user.Name, user.Email, user.PasswordHash, user.RolesToStrings(), user.Status, user.State, user.CreatedAt, user.LastSeenAt, user.IsDeleted, user.DeletedAt).
		PlaceholderFormat(sq.Dollar)

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("%s: %w", op, ErrAlreadyExists) // TODO Сделать константой
			}
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return r.FindById(ctx, user.ID)
}

func (r *UserRepository) Update(ctx context.Context, user *entities.User) (*entities.User, error) {
	const op = "repository.user.Update"

	updateBuilder := sq.Update(usersTable).
		Set("name", user.Name).
		Set("email", user.Email).
		Set("password_hash", user.PasswordHash).
		Set("roles", user.RolesToStrings()).
		Set("status", user.Status).
		Set("state", user.State).
		Set("last_seen_at", user.LastSeenAt).
		Set("is_deleted", user.IsDeleted).
		Set("deleted_at", user.DeletedAt).
		Where(sq.Eq{"id": user.ID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.NoDataFound {
				return nil, fmt.Errorf("%s: %w", op, ErrNotFound) // TODO Сделать константой
			}
			if pgErr.Code == pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("%s: %w", op, ErrAlreadyExists) // TODO Сделать константой + подумать правильная ли ошибка
			}
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return r.FindById(ctx, user.ID)
}

// FIXME у меня тут time.Now(), думаю это плохо?
func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	const op = "repository.user.Delete"

	updateBuilder := sq.Update(usersTable).
		Where(sq.Eq{"id": id}).
		Set("state", "deleted").
		Set("deleted_at", time.Now()).
		PlaceholderFormat(sq.Dollar)

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.NoDataFound {
				return fmt.Errorf("%s: %w", op, ErrNotFound)
			}
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
