CREATE TABLE users
(
  id bigint not null unique,
  url text not null,
  name text not null,
  screen_name text not null unique,
  location text,
  lang text,
  description text,
  background_image text,
  image text,
  banner text,
  statuses_count   bigint not null,
  favourites_count integer not null,
  followers_count integer not null,
  friends_count integer not null,
  created_at timestamp not null,
  updated_at timestamp not null
);
