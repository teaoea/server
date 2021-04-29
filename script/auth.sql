create table auth
(
    id                  bigint                not null
        constraint auth_pk
            primary key,
    hide_article        boolean default false not null,
    hide_article_auth   bigint                not null,
    hide_user           boolean default false not null,
    hide_user_auth      bigint                not null,
    delete_comment      boolean default false not null,
    delete_comment_auth bigint                not null
);

alter table auth
    owner to postgres;

create unique index auth_id_uindex
    on auth (id);

create unique index auth_id_uindex_2
    on auth (id);


