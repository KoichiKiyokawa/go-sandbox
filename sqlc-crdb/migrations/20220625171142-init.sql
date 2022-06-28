
-- +migrate Up
CREATE TABLE accounts (
  id SERIAL PRIMARY KEY,
  username TEXT UNIQUE,
  creat_at DATE DEFAULT CURRENT_TIMESTAMP
);

create table statuses (
  id SERIAL PRIMARY KEY,
  body TEXT,
  created_at DATE DEFAULT CURRENT_TIMESTAMP,
  account_id SERIAL,
  foreign key (account_id) references accounts(id)
);
-- +migrate Down
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS statuses;
