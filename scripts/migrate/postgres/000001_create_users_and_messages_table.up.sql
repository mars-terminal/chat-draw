CREATE TABLE users (
    id bigserial primary key,
    first_name text,
    second_name text,
    nick_name text,
    phone text,
    password text,
    created_at timestamp,
    updated_at timestamp
);

CREATE TABLE messages (
    id bigserial primary key,
    text text,
    chat_id bigint,
    peer_id bigint,
    created_at timestamp,
    updated_at timestamp
);
