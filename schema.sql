create table public.user_accounts
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

alter table public.user_accounts
owner to iresharma;

create index idx_user_accounts_deleted_at
on public.user_accounts (deleted_at);

create table public.auths
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
references public.user_accounts,
metadata_id     text,
settings_id     text
);

alter table public.auths
owner to iresharma;

create index idx_auths_deleted_at
on public.auths (deleted_at);

create table public.sessions
(
id         text not null
primary key,
created_at timestamp with time zone,
updated_at timestamp with time zone,
deleted_at timestamp with time zone,
auth_id    text
constraint fk_sessions_auth
references public.auths
);

alter table public.sessions
owner to iresharma;

create index idx_sessions_deleted_at
on public.sessions (deleted_at);

create table public.user_account_invite_codes
(
id              text not null
primary key,
created_at      timestamp with time zone,
updated_at      timestamp with time zone,
deleted_at      timestamp with time zone,
code            text,
user_account_id text
);

alter table public.user_account_invite_codes
owner to iresharma;

create index idx_user_account_invite_codes_deleted_at
on public.user_account_invite_codes (deleted_at);

create table public.board
(
id varchar(255) not null
primary key
);

alter table public.board
owner to iresharma;

create table public.kanbanlabel
(
id           varchar(255) not null
primary key,
name         varchar(255) not null,
color        varchar(255) not null,
"boardId_id" varchar(255) not null
references public.board
);

alter table public.kanbanlabel
owner to iresharma;

create unique index kanbanlabel_name
on public.kanbanlabel (name);

create index "kanbanlabel_boardId_id"
on public.kanbanlabel ("boardId_id");

create table public.item
(
id       varchar(255) not null
primary key,
status   varchar(255) not null,
title    varchar(255) not null,
"desc"   text         not null,
links    varchar(255) not null,
board_id varchar(255) not null
references public.board,
label_id varchar(255) not null
references public.kanbanlabel
);

alter table public.item
owner to iresharma;

create index item_board_id
on public.item (board_id);

create index item_label_id
on public.item (label_id);

create table public.comment
(
id       varchar(255) not null
primary key,
"userId" varchar(255) not null,
message  text         not null,
item_id  varchar(255) not null
references public.item
);

alter table public.comment
owner to iresharma;

create index comment_item_id
on public.comment (item_id);

create table public.reaction
(
id         varchar(255) not null
primary key,
"userId"   varchar(255) not null,
emoji      varchar(255) not null,
comment_id varchar(255) not null
references public.comment
);

alter table public.reaction
owner to iresharma;

create index reaction_comment_id
on public.reaction (comment_id);

create table public.pages
(
id              text not null
primary key,
created_at      timestamp with time zone,
updated_at      timestamp with time zone,
deleted_at      timestamp with time zone,
route           text,
user_account_id text
);

alter table public.pages
owner to iresharma;

create index idx_pages_deleted_at
on public.pages (deleted_at);

create table public.templates
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
references public.pages,
social          boolean,
social_position text
);

alter table public.templates
owner to iresharma;

create index idx_templates_deleted_at
on public.templates (deleted_at);

create table public.meta
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
references public.templates
);

alter table public.meta
owner to iresharma;

create index idx_meta_deleted_at
on public.meta (deleted_at);

create table public.page_links
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
references public.pages,
icon       text,
social     boolean,
sequence   bigint
);

alter table public.page_links
owner to iresharma;

create index idx_page_links_deleted_at
on public.page_links (deleted_at);

create table public.maillist
(
id         serial
primary key,
email      varchar(255) not null,
topics     varchar(255) not null,
created_at timestamp    not null
);

alter table public.maillist
owner to iresharma;

create unique index maillist_email
on public.maillist (email);

create table public.settings
(
id         text not null
primary key,
created_at timestamp with time zone,
updated_at timestamp with time zone,
deleted_at timestamp with time zone
);

alter table public.settings
owner to iresharma;

create index idx_settings_deleted_at
on public.settings (deleted_at);

create table public.metadata
(
id         text not null
primary key,
created_at timestamp with time zone,
updated_at timestamp with time zone,
deleted_at timestamp with time zone,
name       text,
photo_url  text
);

alter table public.metadata
owner to iresharma;

create index idx_metadata_deleted_at
on public.metadata (deleted_at);

