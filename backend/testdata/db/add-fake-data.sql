INSERT INTO users (first_name, last_name, email) VALUES ('Danny', 'Ochoa', 'danny@ochoa.com');
INSERT INTO users (first_name, last_name, email) VALUES ('LeBron', 'James', 'lebron@james.com');

INSERT INTO timerSettings (owner_id, title) VALUES (1, 'Coding');
INSERT INTO timerSettings (owner_id, title) VALUES (1, 'Music Production');
INSERT INTO timerSettings (owner_id, title) VALUES (1, 'DJing');
INSERT INTO timerSettings (owner_id, title) VALUES (1, 'Piano');
INSERT INTO timerSettings (owner_id, title) VALUES (2, 'Basketball');

-- 1.5 hours in Coding for Danny
INSERT INTO timerSessions (timer_setting_id, session_duration_in_seconds, created_at)
VALUES (1, 3600, '2025-03-15 12:00:00');
INSERT INTO timerSessions (timer_setting_id, session_duration_in_seconds, created_at) 
VALUES (1, 1800, '2025-03-20 10:00:00');

-- 45 minutes in Music Production for Danny
INSERT INTO timerSessions (timer_setting_id, session_duration_in_seconds, created_at)
VALUES (2, 1800, '2025-03-10 6:00:00');
INSERT INTO timerSessions (timer_setting_id, session_duration_in_seconds, created_at)
VALUES (2, 900, '2025-03-12 13:00:00');

-- 10 minutes in DJing for Danny
INSERT INTO timerSessions (timer_setting_id, session_duration_in_seconds, created_at)
VALUES (3, 600, '2025-03-17 22:00:00');

-- No progress yet for Piano for Danny

-- 3 hours of Basketball for LeBron
INSERT INTO timerSessions (timer_setting_id, session_duration_in_seconds, created_at)
VALUES (5, 10800, '2025-03-11 20:00:00');
