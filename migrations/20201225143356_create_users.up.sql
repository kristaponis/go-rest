create table users (
    id bigserial not null primary key,
    email varchar not null unique,
    enc_password varchar not null
);
