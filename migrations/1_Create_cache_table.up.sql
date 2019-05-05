create table gitcrow.cached
(
	pk serial
		constraint cached_pk
			primary key,
	owner char(255) not null,
	repo char(255) not null,
	tag char(255)
);