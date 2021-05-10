create table public."user"
(
    id            bigint                not null
        constraint user_pk
            primary key,
    name          varchar               not null,
    password      varchar               not null,
    email         varchar               not null,
    number        varchar               not null,
    email_active  boolean default false,
    number_active boolean default false,
    ip            varchar,
    created_at    timestamp with time zone,
    is_active     boolean default false,
    is_admin      boolean default false,
    avatar        varchar,
    gender        varchar,
    is_hide       boolean default false not null
);

alter table public."user"
    owner to postgres;

create unique index user_email_uindex
    on public."user" (email);

create unique index user_id_uindex
    on public."user" (id);

create unique index user_name_uindex
    on public."user" (name);

create unique index user_number_uindex
    on public."user" (number);


