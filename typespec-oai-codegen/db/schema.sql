create table users (
  id text primary key,
  name text not null,
  email text not null,
  created_at timestamp not null default CURRENT_TIMESTAMP,
  updated_at timestamp not null default CURRENT_TIMESTAMP
);

create table posts (
  id text primary key,
  title text not null,
  body text not null,
  user_id text not null references users(id),
  created_at timestamp not null default CURRENT_TIMESTAMP,
  updated_at timestamp not null default CURRENT_TIMESTAMP
);