CREATE TABLE accounts (
                          id INT PRIMARY KEY CHECK (id > 0),
                          login VARCHAR (32) NOT NULL,
                          password VARCHAR(32) NOT NULL DEFAULT 123456,
                          token VARCHAR(32) NOT NULL,
                          type INT NOT NULL DEFAULT 1 CHECK (type > 0 AND type < 5)
);

#type 1 - request to register
#type 2 - student
#type 3 - mentor
#type 4 - admin