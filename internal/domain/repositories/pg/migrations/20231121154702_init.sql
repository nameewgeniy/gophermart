-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    uid           UUID      UNIQUE PRIMARY KEY,
    login         VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP
);

CREATE INDEX idx_login ON users (login);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP INDEX idx_login;
-- +goose StatementEnd
