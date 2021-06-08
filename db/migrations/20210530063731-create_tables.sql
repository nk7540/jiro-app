create database if not exists dev_database;
use dev_database;

-- +migrate Up
create table if not exists user(
  id            int          not null auto_increment,
  uid           varchar(255) not null,
  status        varchar(255) not null default 'provisional',
  email         varchar(255) not null,
  nickname      varchar(255) not null default '',
  thumbnail_url text(1024)   not null default '',
  created_at    datetime     not null,
  updated_at    datetime     not null,
  primary key(id)
);

create table if not exists follow(
  id           int          not null auto_increment,
  following_id int          not null,
  follower_id  int          not null,
  created_at   datetime     not null,
  updated_at   datetime     not null,
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
  id         int          not null auto_increment,
  name       varchar(32)  not null,
  created_at datetime     not null,
  updated_at datetime     not null,
  primary key(id)
);

create table if not exists content(
  id            int          not null auto_increment,
  user_id       int          null,
  category_id   int          not null,
  title         varchar(255) not null,
  description   text(1024)   not null default '',
  thumbnail_url text(1024)   not null default '',
  created_at    datetime     not null,
  updated_at    datetime     not null,
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
  id         int          not null auto_increment,
  user_id    int          not null,
  content_id int          not null,
  created_at datetime     not null,
  updated_at datetime     not null,
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

create table if not exists browse(
  id         int          not null auto_increment,
  user_id    int          not null,
  content_id int          not null,
  created_at datetime     not null,
  updated_at datetime     not null,
  primary key(id),
  constraint fk_user_browse
    foreign key(user_id)
    references user (id)
    on delete no action
    on update no action,
  constraint fk_content_browse
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
drop table browse;
set FOREIGN_KEY_CHECKS=1;
