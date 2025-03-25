CREATE TABLE IF NOT EXISTS users (
  id                        BIGSERIAL PRIMARY KEY,
  first_name                VARCHAR(255) NOT NULL,
  last_name                 VARCHAR(255) NOT NULL,
  email                     VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS timerSettings (
  id                        BIGSERIAL PRIMARY KEY,
  owner_id                  BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  title                     VARCHAR(128) NOT NULL
);

CREATE TABLE IF NOT EXISTS timerSessions (
  id                          BIGSERIAL PRIMARY KEY,
  timer_setting_id            BIGINT NOT NULL REFERENCES timerSettings(id) ON DELETE CASCADE,
  session_duration_in_seconds INTEGER NOT NULL,
  created_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
