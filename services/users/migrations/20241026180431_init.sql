-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    roles TEXT[] NOT NULL,
    status VARCHAR(50) NOT NULL,
    state VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    last_seen_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_users_emails ON users (email);

-- dummy admin user
INSERT INTO users (id, name, email, password_hash, roles, status, state, created_at, last_seen_at, deleted_at)
VALUES ('52a82ad8-b265-11ef-874c-0250e2d265f1', 'admin', 'admin@null.ru', 'null_password', '{admin}', 'active', 'active', NOW(), NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS users;
DROP INDEX IF EXISTS idx_users_emails;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
