-- SEED DATA FOR TESTING THE SERVER --
INSERT INTO Users (publicId, name, email) VALUES ('15', 'chris', 'awesome@reallyawesome.com');
INSERT INTO Users (publicId, name, email) VALUES ('666', 'gerron', 'notawesome@nope.com');
INSERT INTO Flags (host, tag, value, comment) VALUES ('192.168.0.220', 'Flag3_Hash1', 1, 'easy flag');
INSERT INTO Flags (host, tag, value, comment) VALUES ('192.168.0.129', 'Flag6_Hash90', 5, 'hard flag');
INSERT INTO UsersFlags (uId, fId) VALUES (1, 1);
INSERT INTO UsersFlags (uId, fId) VALUES (1, 2);