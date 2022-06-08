CREATE TABLE students (
                          id INT PRIMARY KEY AUTO_INCREMENT,
                          fio VARCHAR(60) NOT NULL,
                          keygroup INT NOT NULL,
                          expires TIME NOT NULL,
                          CHECK ( keygroup > 0 )
);

