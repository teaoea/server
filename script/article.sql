create table public.article
(
    id         bigint                not null
        constraint article_pk
            primary key,
    title      varchar               not null,
    content    varchar               not null,
    img        varchar,
    category   varchar               not null,
    show       boolean default false not null,
    view       bigint                not null,
    sha256     varchar               not null,
    author     varchar               not null,
    license    varchar,
    is_hide    boolean default false not null,
    created_at varchar               not null
);

comment on column public.article.content is '存储文件保存的路径';

alter table public.article
    owner to postgres;

create unique index article_body_uindex
    on public.article (content);

create unique index article_id_uindex
    on public.article (id);


