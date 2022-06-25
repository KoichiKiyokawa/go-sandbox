
-- +migrate Up
CREATE TABLE accounts (
  id SERIAL PRIMARY KEY,
  username TEXT UNIQUE,
  creat_at DATE DEFAULT CURRENT_TIMESTAMP
);
-- +migrate Down
DROP TABLE accounts;
