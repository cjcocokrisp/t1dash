CREATE TABLE IF NOT EXISTS users (
  id uuid NOT NULL PRIMARY KEY,
  username varchar(32) UNIQUE NOT NULL,
  first_name varchar(32) NOT NULL,
  last_name varchar(32) NOT NULL,
  password bpchar NOT NULL,
  avatar bpchar NOT NULL DEFAULT 'static/avatars/default.png', role bpchar DEFAULT 'user',
  settings jsonb NOT NULL,
  connections jsonb NOT NULL
);
