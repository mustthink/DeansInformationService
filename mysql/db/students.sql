CREATE TABLE students (
                          id INT PRIMARY KEY,
                          fio VARCHAR(60) NOT NULL,
                          keygroup INT NOT NULL,
                          expires DATE NOT NULL,
                          CHECK ( keygroup > 0 )
);

