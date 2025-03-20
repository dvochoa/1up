CREATE TABLE IF NOT EXISTS users (
  id                        BIGSERIAL PRIMARY KEY,
  first_name                VARCHAR(255) NOT NULL,
  last_name                 VARCHAR(255) NOT NULL,
  email                     VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS timerSettings (
  id                        BIGSERIAL PRIMARY KEY,
  owner_id                  BIGINT REFERENCES users(id),
  title                     VARCHAR(128) NOT NULL
);

CREATE TABLE IF NOT EXISTS timerProgress (
  id                          BIGSERIAL PRIMARY KEY,
  timer_setting_id            BIGINT REFERENCES timerSettings(id),
  session_duration_in_seconds INTEGER NOT NULL,
  session_timestamp           TIMESTAMPTZ NOT NULL
);
