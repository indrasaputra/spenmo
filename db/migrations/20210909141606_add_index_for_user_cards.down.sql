BEGIN;

DROP INDEX IF EXISTS index_on_user_cards_on_user_id_and_deleted_at;
DROP INDEX IF EXISTS index_on_user_cards_on_id_and_user_id_and_deleted_at;

COMMIT;