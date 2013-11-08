DROP TABLE IF EXISTS UsersFlags;
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Flags;

CREATE TABLE Users (
    id INT NOT NULL AUTO_INCREMENT,
    publicId INT NOT NULL,
    name VARCHAR(24) NOT NULL,
    email VARCHAR(32) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE Flags (
    id INT NOT NULL AUTO_INCREMENT,
    host VARCHAR(32) NOT NULL,
    tag VARCHAR(32) NOT NULL,
    value INT NOT NULL,
    discovered BOOLEAN NOT NULL,
    comment VARCHAR(128),
    PRIMARY KEY (id)
);

CREATE TABLE UsersFlags (
    uId INT NOT NULL,
    fId INT NOT NULL,
    FOREIGN KEY (uId) REFERENCES Users (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (fId) REFERENCES Flags (id) ON UPDATE CASCADE ON DELETE CASCADE
);


-- SEED DATA --
-- COMMENT OUT WHEN SHIT GETS REAL --
INSERT INTO Users (publicId, name, email) VALUES (15, 'chris', 'awesome@reallyawesome.com');
INSERT INTO Users (publicId, name, email) VALUES (666, 'gerron', 'notawesome@nope.com');
INSERT INTO Flags (host, tag, value, discovered, comment) VALUES ('192.168.0.220', 'Flag3_Hash1', 1, false, 'easy flag');
INSERT INTO Flags (host, tag, value, discovered, comment) VALUES ('192.168.0.129', 'Flag6_Hash90', 5, false, 'hard flag');
INSERT INTO UsersFlags VALUES (1, 1);
INSERT INTO UsersFlags VALUES (1, 2); 