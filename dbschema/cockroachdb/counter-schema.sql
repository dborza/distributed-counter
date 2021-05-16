CREATE DATABASE IF NOT EXISTS counter_db;

DROP TABLE counter_db.counter;

CREATE TABLE counter_db.counter (
    id INT NOT NULL,
    value INT NOT NULL,
    CONSTRAINT "primary" PRIMARY KEY (id)
  );

USE counter_db;

INSERT INTO counter(id, value) values(1, 0);
