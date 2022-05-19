CREATE TABLE group (
                        keygroup INT PRIMARY KEY CHECK (keygroup > 0),
                        flow INT NOT NULL,
                        numgr INT NOT NULL,
                        course VARCHAR(30) NOT NULL,
                        CHECK (flow > 0 AND numgr > 0)
)
