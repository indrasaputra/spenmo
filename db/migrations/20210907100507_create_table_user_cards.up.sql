BEGIN;

CREATE TABLE IF NOT EXISTS user_cards (
  id            BIGSERIAL       PRIMARY KEY,
  user_id       BIGINT,
  wallet_id     BIGINT,
  limit_daily   DECIMAL,
  limit_monthly DECIMAL,
  created_at    TIMESTAMP,
  updated_at    TIMESTAMP,
  deleted_at    TIMESTAMP,

  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
  CONSTRAINT fk_wallet FOREIGN KEY(wallet_id) REFERENCES user_wallets(id)
);

COMMIT;