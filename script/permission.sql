create table public.permission
(
    id                serial  not null
        constraint permission_pk
            primary key,
    user_id           bigint  not null,
    name              varchar not null,
    hide_article      boolean default false,
    hide_article_auth varchar not null,
    hide_user         boolean default false,
    hide_user_auth    varchar not null,
    del_comment       boolean default false,
    del_comment_auth  varchar not null,
    add_category      boolean default false,
    add_category_auth varchar,
    del_category      boolean default false,
    del_category_auth varchar
);

alter table public.permission
    owner to postgres;

create unique index permission_id_uindex
    on public.permission (id);


