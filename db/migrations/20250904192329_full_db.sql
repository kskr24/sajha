-- +goose Up
-- +goose StatementBegin

-- users table
CREATE TABLE IF NOT EXISTS users
(
    id                   SERIAL PRIMARY KEY,
    username             TEXT UNIQUE NOT NULL,
    email                TEXT NOT NULL UNIQUE,
    name                 TEXT NOT NULL DEFAULT '',  
    password             TEXT         NOT NULL DEFAULT '',
    created_at           TIMESTAMP       NOT NULL DEFAULT now(),
    updated_at           TIMESTAMP       NOT NULL DEFAULT now()
);

-- session table
CREATE TABLE IF NOT EXISTS session
(
    id      BIGSERIAL PRIMARY KEY,
    ua      VARCHAR(255) NOT NULL DEFAULT '', 
    ip      VARCHAR(255) NOT NULL DEFAULT '',
    user_id INT       NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token   VARCHAR(255) NOT NULL DEFAULT '',
    expires TIMESTAMP       NOT NULL DEFAULT now(),
    created TIMESTAMP       NOT NULL DEFAULT now(),
    updated TIMESTAMP       NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_session_token
    ON session (token);

CREATE INDEX IF NOT EXISTS idx_session_user
    ON session (user_id);

-- workspaces table
CREATE TABLE IF NOT EXISTS workspaces(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    owner_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    create_at TIMESTAMP DEFAULT now()
);

-- workspace_members table
CREATE TABLE IF NOT EXISTS workspace_members(
    workspace_id INT NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role TEXT NOT NULL DEFAULT 'member',
    PRIMARY KEY(workspace_id, user_id)
);

-- boards table
CREATE TABLE IF NOT EXISTS boards(
    id SERIAL PRIMARY KEY,
    workspace_id INT NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);


-- tasks table
CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL PRIMARY KEY,
    board_id INT NOT NULL REFERENCES boards(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    details TEXT,
    status TEXT DEFAULT 'todo',
    assignee INT REFERENCES users(id),
    create_at TIMESTAMP DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS workspaces;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS boards;
DROP TABLE IF EXISTS workspace_members;
DROP TABLE IF EXISTS session;
-- +goose StatementEnd
