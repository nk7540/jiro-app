create database if not exists dev_database;
use dev_database;

create table user(
  id                 int          not null auto_increment,
  uid                varchar(255) null,
  nickname           varchar(255) null,
  email              varchar(255) not null,
  encrypted_password varchar(255) not null,
  primary key(id)
);

create table follow(
  id           int not null auto_increment,
  following_id int not null,
  follower_id  int not null,
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

create table category(
  id   int         not null auto_increment,
  name varchar(32) not null,
  primary key(id)
);

create table content(
  id          int          not null auto_increment,
  user_id     int          null,
  category_id int          not null,
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

create table favorite(
  id         int not null auto_increment,
  user_id    int not null,
  content_id int not null,
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
