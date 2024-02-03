CREATE TABLE todos (
  id SERIAL PRIMARY KEY NOT NULL,
  title varchar(150),
  content varchar(500),
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp
);

CREATE INDEX idx_todo_deleted_at ON todos (deleted_at);
