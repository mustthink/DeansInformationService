CREATE TABLE groupinfo (
                        keygroup INT PRIMARY KEY AUTO_INCREMENT,
                        mentorid INT NOT NULL,
                        flow INT NOT NULL,
                        numgr INT NOT NULL,
                        course VARCHAR(30) NOT NULL,
                        CHECK (flow > 0 AND numgr > 0)
);
