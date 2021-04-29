create table category
(
    id   varchar not null
        constraint category_pk
            primary key,
    name varchar not null
);

alter table category
    owner to postgres;

create unique index category_id_uindex
    on category (id);

create unique index category_name_uindex
    on category (name);


