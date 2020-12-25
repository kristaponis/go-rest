create table users (
    id int not null primary key,
    email varchar not null unique,
    enc_password varchar not null
);