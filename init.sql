CREATE DATABASE avito;

\connect avito;

CREATE TABLE
    users(
        id SERIAL NOT NULL PRIMARY KEY,
        name VARCHAR(100),
        balance INT,
        reserve INT
    );