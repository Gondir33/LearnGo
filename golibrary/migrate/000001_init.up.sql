CREATE TABLE IF NOT EXISTS author (
    id SERIAL PRIMARY KEY,
    name_author VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS book (
    id SERIAL PRIMARY KEY,
    name_book VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS author_book (
    id SERIAL PRIMARY KEY,
    id_author INT NOT NULL REFERENCES author (id),
    id_book INT NOT NULL REFERENCES book (id)
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS user_book (
    id SERIAL PRIMARY KEY,
    id_user INT NOT NULL REFERENCES users (id),
    id_book INT UNIQUE NOT NULL
);