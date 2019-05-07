create table users
(
    id int
        constraint users_pk
            primary key,
    username char(255) not null,
    slack_id char(255)
);

create unique index users_username_uindex
    on users (username);