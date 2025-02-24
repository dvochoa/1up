DROP TABLE IF EXISTS timers;
CREATE TABLE timers (
  id         bigserial PRIMARY KEY,
  title      varchar(128) NOT NULL
);
