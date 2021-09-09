BEGIN;

CREATE INDEX IF NOT EXISTS index_on_user_cards_on_user_id_and_deleted_at ON user_cards USING btree (user_id, deleted_at);
CREATE INDEX IF NOT EXISTS index_on_user_cards_on_id_and_user_id_and_deleted_at ON user_cards USING btree (id, user_id, deleted_at);

COMMIT;