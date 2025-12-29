CREATE TABLE IF NOT EXISTS sessions (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id uuid NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  expires_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  last_seen timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  valid boolean NOT NULL,
  ip bpchar NOT NULL
)
