CREATE DATABASE IF NOT EXISTS counter_db;

CREATE TABLE IF NOT EXISTS counter_db.counter (
  id INT NOT NULL,
  value INT NOT NULL,
  CONSTRAINT "primary" PRIMARY KEY (id)
);

USE counter_db;

DELETE FROM counter;

INSERT INTO counter(id, value) values(1, 0);
