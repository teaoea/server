create table public.comment
(
    id      bigint  not null,
    title   bigint  not null,
    "user"  varchar not null,
    content varchar not null,
    time    varchar not null
);

alter table public.comment
    owner to postgres;

create unique index comment_id_uindex
    on public.comment (id);


