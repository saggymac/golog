CREATE TABLE apps (
  id SERIAL,
  name TEXT,
  account_id INTEGER,
  UNIQUE (name, account_id)
);
