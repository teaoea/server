create table permission
(
    id                  serial  not null
        constraint permission_pk
            primary key,
    user_id             bigint  not null,
    name                varchar not null,
    hide_article        boolean default false,
    hide_article_auth   varchar not null,
    hide_user           boolean default false,
    hide_user_auth      varchar not null,
    delete_comment      boolean default false,
    delete_comment_auth varchar not null
);

alter table permission
    owner to postgres;

create unique index permission_id_uindex
    on permission (id);


