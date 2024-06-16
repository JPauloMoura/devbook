SELECT 'CREATE DATABASE devbook' 
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'devbook')\gexec;

USE devbook;

CREATE TABLE IF NOT EXISTS users(
    id serial primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(50) not null,
    createdAt timestamp default current_timestamp
)