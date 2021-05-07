create table public.draft
(
    id         bigint                not null
        constraint draft_pk
            primary key,
    title      varchar               not null,
    body       varchar               not null,
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

comment on column public.draft.body is '存储文件保存的路径';

alter table public.draft
    owner to postgres;

create unique index draft_body_uindex
    on public.draft (body);

create unique index article_id_uindex
    on public.draft (id);