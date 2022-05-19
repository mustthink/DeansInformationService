CREATE TABLE accounts (
                          id INT PRIMARY KEY CHECK (id > 0),
                          login VARCHAR (30) NOT NULL,
                          password VARCHAR(30) NOT NULL DEFAULT 123456,
                          type INT NOT NULL DEFAULT 1 CHECK (type > 0 AND type < 5)
);

