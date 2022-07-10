CREATE TABLE accounts (
                          id INT PRIMARY KEY AUTO_INCREMENT,
                          login VARCHAR (32) NOT NULL,
                          pass VARCHAR(32) NOT NULL DEFAULT 123456,
                          type INT NOT NULL DEFAULT 1 CHECK (type > 0 AND type < 5)
);

#type 1 - student
#type 2 - mentor
#type 3 - admin