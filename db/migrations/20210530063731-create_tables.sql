create database if not exists dev_database;
use dev_database;

-- +migrate Up
create table if not exists user(
  id                 varchar(255) not null,
  uid                varchar(255) not null,
  status             varchar(255) not null,
  email              varchar(255) not null,
  nickname           varchar(255) not null,
  primary key(id)
);

create table if not exists follow(
  id           varchar(255) not null,
  following_id varchar(255) not null,
  follower_id  varchar(255) not null,
  primary key(id),
  constraint fk_user_following
    foreign key(following_id)
    references user (id)
    on delete no action
    on update no action,
  constraint fk_user_follower
    foreign key(follower_id)
    references user (id)
    on delete no action
    on update no action
);

create table if not exists category(
  id   varchar(255) not null,
  name varchar(32)  not null,
  primary key(id)
);

create table if not exists content(
  id          varchar(255) not null,
  user_id     varchar(255) null,
  category_id varchar(255) not null,
  title       varchar(255) not null,
  primary key(id),
  constraint fk_user_content
    foreign key(user_id)
    references user (id)
    on delete no action
    on update no action,
  constraint fk_category_content
    foreign key(category_id)
    references category (id)
    on delete no action
    on update no action
);

create table if not exists favorite(
  id         varchar(255) not null,
  user_id    varchar(255) not null,
  content_id varchar(255) not null,
  primary key(id),
  constraint fk_user_favorite
    foreign key(user_id)
    references user (id)
    on delete no action
    on update no action,
  constraint fk_content_favorite
    foreign key(content_id)
    references content (id)
    on delete no action
    on update no action
);

-- +migrate Down
set FOREIGN_KEY_CHECKS=0;
drop table user;
drop table follow;
drop table category;
drop table content;
drop table favorite;
set FOREIGN_KEY_CHECKS=1;
