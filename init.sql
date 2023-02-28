CREATE DATABASE avito;

\connect avito;

CREATE TABLE
    users(
        id SERIAL NOT NULL PRIMARY KEY,
        name VARCHAR(100),
        balance INT,
        reserve INT
    );
CREATE TABLE
    orders (
        id SERIAL NOT NULL PRIMARY KEY,
        userid INT,
        taksid INT,
        cost INT
    );
CREATE TABLE
    finances (
        id SERIAL NOT NULL PRIMARY KEY,
        userid INT,
        taksid INT,
        cost INT
    );