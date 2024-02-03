CREATE TABLE users (
  id SERIAL PRIMARY KEY NOT NULL,
  name varchar(150),
  avatar varchar(200),
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp
);

CREATE INDEX idx_users_deleted_at ON users (deleted_at);