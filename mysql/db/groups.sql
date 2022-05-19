CREATE TABLE groupinfo (
                        keygroup INT PRIMARY KEY CHECK (keygroup > 0),
                        mentorid INT NOT NULL,
                        flow INT NOT NULL,
                        numgr INT NOT NULL,
                        course VARCHAR(30) NOT NULL,
                        CHECK (flow > 0 AND numgr > 0)
);
