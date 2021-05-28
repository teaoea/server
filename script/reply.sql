create table public.reply
(
    id         bigint  not null
        constraint comment_two_pk
            primary key,
    comment    bigint  not null,
    "user"     varchar not null,
    content    varchar not null,
    created_at varchar not null
);

alter table public.reply
    owner to postgres;

create unique index comment_two_id_uindex
    on public.reply (id);


