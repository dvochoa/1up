CREATE TABLE IF NOT EXISTS timers (
  id         bigserial PRIMARY KEY,
  title      varchar(128) NOT NULL
);

INSERT INTO timers(title) VALUES ('Coding');
INSERT INTO timers(title) VALUES ('Music Production');
INSERT INTO timers(title) VALUES ('DJing');
INSERT INTO timers(title) VALUES ('Piano');
