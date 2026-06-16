-- +goose Up
CREATE TABLE users(
   user_id INTEGER PRIMARY KEY AUTOINCREMENT,
   username TEXT
);

-- +goose Down
DROP TABLE users;
