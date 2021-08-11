create table article
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

comment on column article.content is '存储文件保存的路径';

alter table article
    owner to postgres;

create unique index article_body_uindex
    on article (content);

create unique index article_id_uindex
    on article (id);

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

create table comment
(
    id         bigint  not null
        constraint comment_pk
            primary key,
    title      bigint  not null,
    "user"     varchar not null,
    content    varchar not null,
    created_at varchar not null
);

alter table comment
    owner to postgres;

create unique index comment_id_uindex
    on comment (id);

create table draft
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

comment on column draft.body is '存储文件保存的路径';

alter table draft
    owner to postgres;

create unique index draft_body_uindex
    on draft (body);

create table github
(
    id        bigint  not null
        constraint github_pk
            primary key,
    github_id varchar not null,
    name      varchar not null,
    email     varchar not null
);

alter table github
    owner to postgres;

create unique index github_github_id_uindex
    on github (github_id);

create unique index github_id_uindex
    on github (id);

create unique index github_name_uindex
    on github (name);

create table permission
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

alter table permission
    owner to postgres;

create unique index permission_id_uindex
    on permission (id);

create table reply
(
    id         bigint  not null
        constraint comment_two_pk
            primary key,
    comment    bigint  not null,
    "user"     varchar not null,
    content    varchar not null,
    created_at varchar not null
);

alter table reply
    owner to postgres;

create unique index comment_two_id_uindex
    on reply (id);

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


