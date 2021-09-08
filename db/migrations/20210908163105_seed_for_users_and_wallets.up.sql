BEGIN;

-- insert 5 users
INSERT INTO users (name, email, password, created_at, updated_at) VALUES ('a', 'a@a', 'a', NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO users (name, email, password, created_at, updated_at) VALUES ('b', 'b@b', 'b', NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO users (name, email, password, created_at, updated_at) VALUES ('c', 'c@c', 'c', NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO users (name, email, password, created_at, updated_at) VALUES ('d', 'd@d', 'd', NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO users (name, email, password, created_at, updated_at) VALUES ('e', 'e@e', 'e', NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- insert 5 wallets for user 1
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (1, 10000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (1, 20000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (1, 30000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (1, 40000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (1, 50000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- insert 5 wallets for user 2
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (2, 10000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (2, 20000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (2, 30000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (2, 40000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (2, 50000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- insert 5 wallets for user 3
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (3, 10000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (3, 20000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (3, 30000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (3, 40000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (3, 50000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- insert 5 wallets for user 4
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (4, 10000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (4, 20000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (4, 30000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (4, 40000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (4, 50000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

-- insert 5 wallets for user 5
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (5, 10000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (5, 20000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (5, 30000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (5, 40000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);
INSERT INTO user_wallets (user_id, balance, created_at, updated_at) VALUES (5, 50000000, NOW()::TIMESTAMP, NOW()::TIMESTAMP);

COMMIT;