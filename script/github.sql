create table github
(
    id        bigint  not null
        constraint github_pk
            primary key,
    github_id varchar not null,
    name      varchar not null,
    email     varchar not null
);

alter table github
    owner to postgres;

create unique index github_github_id_uindex
    on github (github_id);

create unique index github_id_uindex
    on github (id);

create unique index github_name_uindex
    on github (name);


