CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
  birthday DATE NOT NULL
);

INSERT INTO users
VALUES (1, 'hoge');
