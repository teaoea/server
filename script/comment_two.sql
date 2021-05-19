create table public.comment_two
(
    id      bigint  not null
        constraint comment_two_pk
            primary key,
    comment bigint  not null,
    "user"  varchar not null,
    content varchar not null,
    time    varchar not null
);

alter table public.comment_two
    owner to postgres;

create unique index comment_two_id_uindex
    on public.comment_two (id);


