CREATE DATABASE accountdb;

USE accountdb;

CREATE TABLE accounts (
                          id INT auto_increment PRIMARY KEY CHECK (id > 0),
                          login VARCHAR (30) NOT NULL,
                          password VARCHAR(30) NOT NULL DEFAULT 123456,
                          type INT NOT NULL DEFAULT 3 CHECK (type > 0 AND type < 4)
)

