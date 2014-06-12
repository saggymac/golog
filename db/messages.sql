CREATE TABLE messages (
  id BIGSERIAL,
  account_id INTEGER,
  app_id INTEGER,
  ts TIMESTAMP WITH TIME ZONE,
  msg TEXT
);
