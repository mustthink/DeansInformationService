CREATE TABLE mentors (
                           id INT PRIMARY KEY CHECK (id > 0),
                           name VARCHAR(60) NOT NULL,
                           chair VARCHAR(30)
);
