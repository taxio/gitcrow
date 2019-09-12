create table repo
(
    id int not null
        constraint repo_pk
            primary key autoincrement,
    host text not null,
    owner text not null,
    name text not null
);

create unique index repo_host_owner_name_uindex
    on repo (host, owner, name);