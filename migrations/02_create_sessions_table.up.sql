CREATE TABLE IF NOT EXISTS sessions (
  id uuid NOT NULL PRIMARY KEY,
  user_id uuid NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  expires_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  ip inet NOT NULL,
  user_agent bpchar DEFAULT 'UNKNOWN',
  last_access timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
)
