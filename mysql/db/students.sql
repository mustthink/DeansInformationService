CREATE TABLE students (
                          id INT PRIMARY KEY CHECK (id > 0),
                          name VARCHAR(60) NOT NULL,
                          keygroup INT
)

