CREATE TABLE tracks
(
    track_id serial    primary key,
    author varchar(70)  not null,
    title  varchar(70)  not null,
    genre varchar(70)  not null,
    album varchar(70)  not null,
    duration int not null
);

CREATE TABLE users
(
    user_id     serial    primary key,
    first_name varchar(70)  not null,
    last_name  varchar(70)  not null,
    email      varchar(255) not null,
    password   varchar(255) not null
);

CREATE TABLE tracklist
(
    id serial primary key,
    track_id int references tracks (track_id) on delete cascade not null,
    user_id int references users (user_id) on delete cascade  not null
);