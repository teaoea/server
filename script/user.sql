create table "user"
(
    id           bigint                not null
        constraint user_pk
            primary key,
    username     varchar               not null,
    password     varchar               not null,
    email        varchar               not null,
    phone        varchar               not null,
    email_active boolean default false,
    phone_active boolean default false,
    created_at   timestamp with time zone,
    is_active    boolean default false,
    is_admin     boolean default false,
    avatar       varchar,
    gender       varchar,
    is_hide      boolean default false not null,
    prefix       varchar               not null
);

alter table "user"
    owner to postgres;

create unique index user_email_uindex
    on "user" (email);

create unique index user_id_uindex
    on "user" (id);

create unique index user_name_uindex
    on "user" (username);

create unique index user_number_uindex
    on "user" (phone);


