CREATE TABLE IF NOT EXISTS CUSTOMERS (
    ID BIGINT PRIMARY KEY,
    NAME VARCHAR(255) NOT NULL,
    AGE INT
);

INSERT INTO CUSTOMERS(ID, NAME, AGE)
VALUES
(1, 'Harry', 27),
(2, 'Ginny', 22),
(3, 'Ron', 27),
(4, 'Hermoine', 27),
(5, 'Lupin', 48),
(6, 'Hagrid', 89);
