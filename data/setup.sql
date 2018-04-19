
#-- chitchat 数据库创建脚本
create database if not exists chitchat default charset utf8 collate utf8_general_ci;
use chitchat;
#-------------------------------------------------------------------------------------
drop table if exists posts;
drop table if exists threads;
drop table if exists sessions;
#-------------------------------------------------------------------------------------
-- users
drop table if exists users;
create table users(
  `id`          int(4) auto_increment,
  `uuid`        varchar(64)  not null,
  `name`        varchar(255),
  `email`       varchar(255) not null,
  `password`    varchar(255) not null,
  `created_at`  timestamp default current_timestamp,

  constraint pk_users primary key(`id`),
  constraint uk_users_uuid unique key(`uuid`),
  constraint uk_users_email unique key(`email`)
) engine=InnoDB default charset=utf8;

-- sessions
drop table if exists sessions;
create table sessions(
  `id`          int(4) auto_increment,
  `uuid`        varchar(64)  not null,
  `email`       varchar(255),
  `user_id`     int(4)       not null,
  `created_at`  timestamp default current_timestamp,

  constraint pk_sessions primary key(`id`),
  constraint uk_sessions_uuid unique key(`uuid`),
  constraint fk_sessions_user_id foreign key(`user_id`) references users(`id`)
) engine=InnoDB default charset=utf8;

-- threads
drop table if exists threads;
create table threads(
  `id`          int(4) auto_increment,
  `uuid`        varchar(64)  not null,
  `topic`       text,
  `user_id`     int(4)       not null,
  `created_at`  timestamp default current_timestamp,

  constraint pk_threads primary key(`id`),
  constraint uk_threads_uuid unique key(`uuid`),
  constraint fk_threads_user_id foreign key(`user_id`) references users(`id`)
) engine=InnoDB default charset=utf8;

-- posts
drop table if exists posts;
create table posts(
  `id`          int(4) auto_increment,
  `uuid`        varchar(64)  not null,
  `body`        text,
  `user_id`     int(4)       not null,
  `thread_id`   int(4)       not null,
  `created_at`  timestamp default current_timestamp,

  constraint pk_posts primary key(`id`),
  constraint uk_posts_uuid unique key(`uuid`),
  constraint fk_posts_user_id foreign key(`user_id`) references users(`id`),
  constraint fk_posts_thread_id foreign key(`thread_id`) references threads(`id`)
) engine=InnoDB default charset=utf8;

#-------------------------------------------------------------------------------------
