-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    password TEXT NOT NULL DEFAULT ''

);

-- session table
CREATE TABLE IF NOT EXISTS session
(
    id      BIGSERIAL PRIMARY KEY,
    ua      VARCHAR(255) NOT NULL DEFAULT '', -- User agent information
    ip      VARCHAR(255) NOT NULL DEFAULT '', -- IP address
    user_id BIGINT       NOT NULL,            -- Foreign key to user table
    token   VARCHAR(255) NOT NULL DEFAULT '', -- Session token
    expires BIGINT       NOT NULL DEFAULT 0,  -- Expiration timestamp
    created BIGINT       NOT NULL DEFAULT 0,  -- Creation timestamp
    updated BIGINT       NOT NULL DEFAULT 0   -- Last updated timestamp
);

CREATE INDEX IF NOT EXISTS idx_session_token
    ON session (token);

CREATE INDEX IF NOT EXISTS idx_session_user
    ON session (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS session;
-- +goose StatementEnd
