-- drop table users;
-- drop table sessions;
-- drop table leaderboard;
-- drop table question_table;

create table if not exists users 
(
  id serial primary key,
  uuid varchar(64) not null unique,
  username varchar(255) not null unique,
  email varchar(255) not null unique,
  password varchar(255) not null,
  created_at timestamp not null
);

create table if not exists sessions
(
  id serial primary key,
  uuid varchar(64) not null unique,
  email varchar(255),
  user_id integer references users(id),
  created_at timestamp not null
);

create table if not exists leaderboard
(
  id serial primary key,
  username varchar(255) references users(username),
  level integer DEFAULT 0,
  solve_time timestamp not null
);

create table if not exists question_table
(
  id serial primary key,
  question varchar(100) not null
);