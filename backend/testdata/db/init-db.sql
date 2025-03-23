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

CREATE TABLE IF NOT EXISTS timerProgress (
  id                          BIGSERIAL PRIMARY KEY,
  timer_setting_id            BIGINT NOT NULL REFERENCES timerSettings(id) ON DELETE CASCADE,
  session_duration_in_seconds INTEGER NOT NULL,
  session_timestamp           TIMESTAMPTZ NOT NULL
);

-- TODO: tests should just reference the real init scripts rather than duplicating schema

INSERT INTO users (first_name, last_name, email) VALUES ('Danny', 'Ochoa', 'danny@ochoa.com');
INSERT INTO users (first_name, last_name, email) VALUES ('LeBron', 'James', 'lebron@james.com');

INSERT INTO timerSettings (owner_id, title) VALUES (1, 'Coding');
INSERT INTO timerSettings (owner_id, title) VALUES (1, 'Music Production');
INSERT INTO timerSettings (owner_id, title) VALUES (1, 'DJing');
INSERT INTO timerSettings (owner_id, title) VALUES (1, 'Piano');
INSERT INTO timerSettings (owner_id, title) VALUES (2, 'Basketball');

-- 1.5 hours in Coding for Danny
INSERT INTO timerProgress (timer_setting_id, session_duration_in_seconds, session_timestamp)
VALUES (1, 3600, '2025-03-15 12:00:00');
INSERT INTO timerProgress (timer_setting_id, session_duration_in_seconds, session_timestamp) 
VALUES (1, 1800, '2025-03-20 10:00:00');

-- 45 minutes in Music Production for Danny
INSERT INTO timerProgress (timer_setting_id, session_duration_in_seconds, session_timestamp)
VALUES (2, 1800, '2025-03-10 6:00:00');
INSERT INTO timerProgress (timer_setting_id, session_duration_in_seconds, session_timestamp)
VALUES (2, 900, '2025-03-12 13:00:00');

-- 10 minutes in DJing for Danny
INSERT INTO timerProgress (timer_setting_id, session_duration_in_seconds, session_timestamp)
VALUES (3, 600, '2025-03-17 22:00:00');

-- No progress yet for Piano for Danny

-- 3 hours of Basketball for LeBron
INSERT INTO timerProgress (timer_setting_id, session_duration_in_seconds, session_timestamp)
VALUES (5, 10800, '2025-03-11 20:00:00');
