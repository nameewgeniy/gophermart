-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id            UUID UNIQUE PRIMARY KEY,
    login         VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255)        NOT NULL,
    balance       BIGINT,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP
);

CREATE INDEX idx_login ON users (login);

CREATE TABLE orders
(
    id          UUID PRIMARY KEY,
    user_id     UUID UNIQUE,
    number      VARCHAR(255) UNIQUE,
    status      VARCHAR(255),
    accrual     BIGINT,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE INDEX idx_number ON orders (number);
CREATE INDEX idx_user_number ON orders (user_id, number);

CREATE TABLE transactions
(
    id           UUID PRIMARY KEY,
    user_id      UUID UNIQUE,
    "order"      VARCHAR(255) UNIQUE,
    sum          BIGINT,
    processed_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE INDEX idx_order ON transactions ("order");
CREATE INDEX idx_user_order ON transactions (user_id, "order");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS orders;
DROP INDEX idx_login;
DROP INDEX idx_number;
DROP INDEX idx_user_number;
-- +goose StatementEnd
