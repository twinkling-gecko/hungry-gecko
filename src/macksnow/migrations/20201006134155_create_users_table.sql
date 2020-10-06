-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id INT PRIMARY KEY AUTO_INCREMENT,
    nickname VARCHAR(15),
    screen_name VARCHAR(15) UNIQUE,
    email VARCHAR(254),
    password_hash TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    confirmation_token TEXT,
    confirmed_at DATETIME
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
