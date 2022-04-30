CREATE TABLE users (
    id bigserial not null,
    first_name text not null ,
    second_name text not null,
    nick_name text not null ,
    email text not null ,
    phone text not null,
    password text not null ,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

CREATE TABLE messages (
    id bigserial not null ,
    text text not null ,
    chat_id bigint,
    peer_id bigint,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

