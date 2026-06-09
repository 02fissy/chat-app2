-- +goose Up
CREATE TABLE IF NOT EXISTS users(
   user_id INTEGER PRIMARY KEY AUTOINCREMENT,
   username TEXT
);

-- +goose Down
DROP TABLE users;
