-- +goose Up
create table chats (
    id serial primary key,
    created_at timestamp not null default now()

);

create table chat_members (
    chat_id int references chats(id) on delete cascade,
    username varchar(255) not null,
    primary key (chat_id, username)
);

create table messages (
    id serial primary key,
    chat_id int references chats(id) on delete cascade ,
    from_username varchar(255) not null,
    text text not null,
    timestamp timestamp not null  default now()

);

-- +goose Down
drop table if exists chats;
drop table if exists chat_members;
drop table if exists messages;