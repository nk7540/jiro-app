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
  is_close     tinyint(1)   not null default 0,
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

create table if not exists comment(
  id         int          not null auto_increment,
  user_id    int          not null,
  content_id int          not null,
  body       text(1024)   not null,
  created_at datetime     not null,
  updated_at datetime     not null,
  primary key(id)
);

create table if not exists favorite(
  id             int          not null auto_increment,
  user_id        int          not null,
  content_id     int          not null,
  comment_id     int          null,
  to_user_ids    varchar(255) not null default '',
  to_close_users tinyint(1)   not null default 0,
  created_at     datetime     not null,
  updated_at     datetime     not null,
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
    on update no action,
  constraint fk_comment_favorite
    foreign key(comment_id)
    references comment (id)
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

create table if not exists notice(
  id            int          not null primary key auto_increment,
  user_id       int          not null,
  type          int          not null,
  is_read       tinyint(1)   not null default 0,
  created_at    datetime     not null,
  updated_at    datetime     not null,
  constraint fk_user_notice foreign key(user_id) references user (id)
);

create table if not exists notice_favorite(
  notice_id             int          not null primary key,
  favorite_id           int          not null,
  user_id               int          not null,
  user_thumbnail_url    text(1024)   not null,
  header                varchar(255) not null,
  body                  varchar(255) not null,
  content_id            int          not null,
  content_thumbnail_url text(1024)   not null,
  created_at            datetime     not null,
  updated_at            datetime     not null,
  constraint fk_notice_notice_favorite foreign key(notice_id) references notice (id),
  constraint fk_favorite_notice_favorite foreign key(favorite_id) references favorite (id),
  constraint fk_user_notice_favorite foreign key(user_id) references user (id),
  constraint fk_content_notice_favorite foreign key(content_id) references content (id)
);

create table if not exists notice_followed(
  notice_id          int          not null primary key,
  user_id            int          not null,
  user_thumbnail_url text(1024)   not null,
  body               varchar(255) not null,
  created_at         datetime     not null,
  updated_at         datetime     not null,
  constraint fk_notice_notice_followed foreign key(notice_id) references notice (id),
  constraint fk_user_notice_followed foreign key(user_id) references user (id)
);

-- +migrate Down
set FOREIGN_KEY_CHECKS=0;
drop table if exists user;
drop table if exists follow;
drop table if exists category;
drop table if exists content;
drop table if exists comment;
drop table if exists favorite;
drop table if exists browse;
drop table if exists notice;
drop table if exists notice_favorite;
drop table if exists notice_followed;
set FOREIGN_KEY_CHECKS=1;
