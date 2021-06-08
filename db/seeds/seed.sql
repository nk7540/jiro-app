insert into user
  (uid, email, nickname)
values
  ('', 'info+user1@blail.co.jp', 'タロウ'),
  ('', 'info+user2@blail.co.jp', 'ハナコ');

insert into follow
  (following_id, follower_id)
values
  (1, 2);

insert into category
  (name)
values
  ('video'),
  ('music'),
  ('comic'),
  ('picture');

insert into content
  (category_id, title)
values
  (1, '動画1'),
  (2, '音楽1');

insert into favorite
  (user_id, content_id)
values
  (1, 1),
  (1, 2);
