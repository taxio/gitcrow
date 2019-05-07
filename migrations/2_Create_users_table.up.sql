create table users
(
    id serial not null,
    name char(255) not null,
    slack_id char(255)
);

create unique index users_id_uindex
    on users (id);

create unique index users_name_uindex
    on users (name);

alter table users
    add constraint users_pk
        primary key (id);

