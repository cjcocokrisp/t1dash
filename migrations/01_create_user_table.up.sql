    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  username varchar(32) UNIQUE NOT NULL,
  firstname varchar(32) NOT NULL,
  lastname varchar(32) NOT NULL,
  password bpchar NOT NULL,
  avatar bpchar NOT NULL DEFAULT 'static/avatars/default.png',
  role bpchar DEFAULT 'user',
  settings jsonb NOT NULL,
  connections jsonb NOT NULL
);
