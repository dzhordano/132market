-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email VARCHAR(100) NOT NULL UNIQUE,
    roles TEXT[] NOT NULL,
    state VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_users_emails ON users (email);

CREATE TABLE IF NOT EXISTS roles_permissions(
    role TEXT NOT NULL,
    permission VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles_permissions;

DROP TYPE IF EXISTS ROLE_TYPE;
DROP TYPE IF EXISTS STATE_TYPE;
-- +goose StatementEnd
