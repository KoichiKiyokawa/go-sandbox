CREATE TABLE users (
  id TEXT PRIMARY KEY,
  name TEXT,
  email TEXT
);

CREATE TABLE posts (
  id TEXT PRIMARY KEY,
  user_id TEXT REFERENCES users(id),
  title TEXT,
  body TEXT
);
