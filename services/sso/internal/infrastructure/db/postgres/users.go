package postgres

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
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
	insertQuery := sq.Insert("users").
		Columns("id", "email", "roles", "state", "created_at", "updated_at", "deleted_at").
		Values(user.ID, user.Email, user.Roles, user.State, user.CreatedAt, user.UpdatedAt, user.DeletedAt).
		PlaceholderFormat(sq.Dollar)

	query, args, err := insertQuery.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	return err
}

func (r *UsersRepository) FindById(ctx context.Context, id string) (*entities.User, error) {
	selectQuery := sq.Select("id", "email", "roles", "state", "created_at", "updated_at", "deleted_at").
		From("users").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectQuery.ToSql()
	if err != nil {
		return nil, err
	}

	var user entities.User
	if err := r.db.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Email, &user.Roles, &user.State, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UsersRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	selectQuery := sq.Select("id", "email", "roles", "state", "created_at", "updated_at", "deleted_at").
		From("users").
		Where(sq.Eq{"email": email}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectQuery.ToSql()
	if err != nil {
		return nil, err
	}

	var user entities.User
	if err := r.db.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Email, &user.Roles, &user.State, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
		return nil, err
	}

	return &user, nil

}

func (r *UsersRepository) Update(ctx context.Context, user *entities.User) error {
	updateQuery := sq.Update("users").
		Set("email", user.Email).
		Set("roles", user.Roles).
		Set("state", user.State).
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": user.ID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := updateQuery.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	return err
}

// SOFT-DELETION METHOD
func (r *UsersRepository) Delete(ctx context.Context, id string) error {
	updateQuery := sq.Update("users").
		Set("state", "deleted").
		Set("deleted_at", time.Now().UTC()).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := updateQuery.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	return err
}

func (r *UsersRepository) AssignRoleById(ctx context.Context, id string, role string) error {
	updateQuery := sq.Update("users").
		Set("roles", sq.Expr("array_append(roles, ?)", role)).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := updateQuery.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	return err
}

func (r *UsersRepository) RevokeRoleById(ctx context.Context, id string, role string) error {
	updateQuery := sq.Update("users").
		Set("roles", sq.Expr("array_remove(roles, ?)", role)).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := updateQuery.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	return err
}
