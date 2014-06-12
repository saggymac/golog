CREATE TABLE accounts (
  account_id INTEGER,
  user_id INTEGER,
  UNIQUE (account_id, user_id)
);
