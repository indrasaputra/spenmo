BEGIN;

CREATE TABLE IF NOT EXISTS user_wallets (
  id            BIGSERIAL       PRIMARY KEY,
  user_id       BIGINT,
  balance       DECIMAL,
  created_at    TIMESTAMP,
  updated_at    TIMESTAMP,
  deleted_at    TIMESTAMP,

  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);

COMMIT;