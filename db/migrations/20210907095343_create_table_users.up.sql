BEGIN;

CREATE TABLE IF NOT EXISTS users (
  id            BIGSERIAL       PRIMARY KEY,
  name          VARCHAR(255),
  email         VARCHAR(255)    UNIQUE NOT NULL,
  password      VARCHAR(255)    NOT NULL,
  created_at    TIMESTAMP,
  updated_at    TIMESTAMP,
  deleted_at    TIMESTAMP
);

COMMIT;