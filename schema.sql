create table if not exists user_accounts
(
    id           text not null
    primary key,
    created_at   timestamp with time zone,
    updated_at   timestamp with time zone,
    deleted_at   timestamp with time zone,
    account_name text,
    email        text,
    photo_url    text,
    owner        text,
    page_id      text,
    bucket_id    text,
    board_id     text
);

alter table user_accounts
    owner to iresharma;

create index if not exists idx_user_accounts_deleted_at
    on user_accounts (deleted_at);

create table if not exists auths
(
    id              text not null
    primary key,
    created_at      timestamp with time zone,
    updated_at      timestamp with time zone,
    deleted_at      timestamp with time zone,
    email           text
    unique,
    password_hash   text,
    salt            text,
    perm            text,
    user_account_id text
    constraint fk_user_accounts_users
    references user_accounts,
    metadata_id     text,
    settings_id     text
);

alter table auths
    owner to iresharma;

create index if not exists idx_auths_deleted_at
    on auths (deleted_at);

create table if not exists sessions
(
    id         text not null
    primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    auth_id    text
    constraint fk_sessions_auth
    references auths
);

alter table sessions
    owner to iresharma;

create index if not exists idx_sessions_deleted_at
    on sessions (deleted_at);

create table if not exists user_account_invite_codes
(
    id              text not null
    primary key,
    created_at      timestamp with time zone,
    updated_at      timestamp with time zone,
    deleted_at      timestamp with time zone,
    code            text,
    user_account_id text
);

alter table user_account_invite_codes
    owner to iresharma;

create index if not exists idx_user_account_invite_codes_deleted_at
    on user_account_invite_codes (deleted_at);

create table if not exists board
(
    id varchar(255) not null
    primary key
    );

alter table board
    owner to iresharma;

create table if not exists kanbanlabel
(
    id           varchar(255) not null
    primary key,
    name         varchar(255) not null,
    color        varchar(255) not null,
    "boardId_id" varchar(255) not null
    references board
    );

alter table kanbanlabel
    owner to iresharma;

create unique index if not exists kanbanlabel_name
    on kanbanlabel (name);

create index if not exists "kanbanlabel_boardId_id"
    on kanbanlabel ("boardId_id");

create table if not exists item
(
    id       varchar(255) not null
    primary key,
    status   varchar(255) not null,
    title    varchar(255) not null,
    "desc"   text         not null,
    links    varchar(255) not null,
    board_id varchar(255) not null
    references board,
    label_id varchar(255) not null
    references kanbanlabel
    );

alter table item
    owner to iresharma;

create index if not exists item_board_id
    on item (board_id);

create index if not exists item_label_id
    on item (label_id);

create table if not exists comment
(
    id       varchar(255) not null
    primary key,
    "userId" varchar(255) not null,
    message  text         not null,
    item_id  varchar(255) not null
    references item
    );

alter table comment
    owner to iresharma;

create index if not exists comment_item_id
    on comment (item_id);

create table if not exists reaction
(
    id         varchar(255) not null
    primary key,
    "userId"   varchar(255) not null,
    emoji      varchar(255) not null,
    comment_id varchar(255) not null
    references comment
    );

alter table reaction
    owner to iresharma;

create index if not exists reaction_comment_id
    on reaction (comment_id);

create table if not exists pages
(
    id              text not null
    primary key,
    created_at      timestamp with time zone,
    updated_at      timestamp with time zone,
    deleted_at      timestamp with time zone,
    route           text,
    user_account_id text
);

alter table pages
    owner to iresharma;

create index if not exists idx_pages_deleted_at
    on pages (deleted_at);

create table if not exists templates
(
    id              text not null
    primary key,
    created_at      timestamp with time zone,
    updated_at      timestamp with time zone,
    deleted_at      timestamp with time zone,
    name            text,
    "desc"          text,
    image           text,
    button          text,
    background      text,
    font            text,
    font_color      text,
    page_id         text
    constraint fk_pages_template
    references pages,
    social          boolean,
    social_position text
);

alter table templates
    owner to iresharma;

create index if not exists idx_templates_deleted_at
    on templates (deleted_at);

create table if not exists meta
(
    id          text not null
    primary key,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone,
    value       text,
    type        text,
    template_id text
    constraint fk_templates_meta_tags
    references templates
);

alter table meta
    owner to iresharma;

create index if not exists idx_meta_deleted_at
    on meta (deleted_at);

create table if not exists page_links
(
    id         text not null
    primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name       text,
    link       text,
    page_id    text
    constraint fk_pages_links
    references pages,
    icon       text,
    social     boolean,
    sequence   bigint
);

alter table page_links
    owner to iresharma;

create index if not exists idx_page_links_deleted_at
    on page_links (deleted_at);

create table if not exists maillist
(
    id         serial
    primary key,
    email      varchar(255) not null,
    topics     varchar(255) not null,
    created_at timestamp    not null
    );

alter table maillist
    owner to iresharma;

create unique index if not exists maillist_email
    on maillist (email);

create table if not exists settings
(
    id         text not null
    primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);

alter table settings
    owner to iresharma;

create index if not exists idx_settings_deleted_at
    on settings (deleted_at);

create table if not exists metadata
(
    id         text not null
    primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name       text,
    photo_url  text
);

alter table metadata
    owner to iresharma;

create index if not exists idx_metadata_deleted_at
    on metadata (deleted_at);

