CREATE TABLE IF NOT EXISTS users
(
     id         SERIAL PRIMARY KEY,
     uuid       VARCHAR
(64) NOT NULL UNIQUE,
     username   VARCHAR
(255) NOT NULL UNIQUE,
     email      VARCHAR
(255) NOT NULL UNIQUE,
     password   VARCHAR
(255) NOT NULL,
     created_at TIMESTAMP NOT NULL
  );

CREATE TABLE IF NOT EXISTS sessions
(
     id         SERIAL PRIMARY KEY,
     uuid       VARCHAR (64) NOT NULL UNIQUE,
     email      VARCHAR
(255),
     user_id    INTEGER REFERENCES users
(id),
     created_at TIMESTAMP NOT NULL
  );

CREATE TABLE IF NOT EXISTS leaderboard
(
     id         SERIAL PRIMARY KEY,
     username   VARCHAR
(255) REFERENCES users
(username),
     level      INTEGER DEFAULT 0,
     solve_time TIMESTAMP NOT NULL
  );


CREATE TABLE IF NOT EXISTS question_table
(
     id       SERIAL PRIMARY KEY,
     question VARCHAR(100) NOT NULL,
     answer VARCHAR(100) NOT NULL
  );


  CREATE TABLE IF NOT EXISTS answer_log
  (
       id SERIAL PRIMARY KEY,
       username VARCHAR(100) NOT NULL,
       level INT NOT NULL,
       attempt VARCHAR NOT NULL,
       at_time TIMESTAMP NOT NULL
  );