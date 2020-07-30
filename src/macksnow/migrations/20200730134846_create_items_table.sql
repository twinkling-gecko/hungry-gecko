-- +goose Up
-- +goose StatementBegin
CREATE TABLE items (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name TEXT,
    summary TEXT,
    uri TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE items;
-- +goose StatementEnd
