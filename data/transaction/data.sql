DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_mode') THEN
        CREATE TYPE TRANSACTION_MODE AS ENUM ('CHEQUE', 'NEFT', 'IMPS', 'RTGS', 'UPI');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS TRANSACTIONS (
    ID SERIAL PRIMARY KEY,
    CUSTOMER_ID INT NOT NULL,
    AMOUNT BIGINT NOT NULL,
    CREATED_AT TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    MODE TRANSACTION_MODE NOT NULL
);


INSERT INTO TRANSACTIONS(CUSTOMER_ID, AMOUNT, CREATED_AT, MODE)
VALUES
-- Harry's Transactions
(1, 100000, NOW(), 'NEFT'),
(1, 100000, NOW(), 'IMPS'),
(1, 50000, NOW(), 'UPI'),
(1, 100000, NOW(), 'NEFT'),
(1, 100000, NOW(), 'NEFT'),
(1, 100000, NOW(), 'NEFT'),
(1, 100000, NOW(), 'NEFT'),
-- Ron's Transactions
(3, 100000, NOW(), 'NEFT'),
(3, 100000, NOW(), 'NEFT'),
(3, -1000, NOW(), 'UPI'),
(3, -1000, NOW(), 'UPI'),
(3, -1000, NOW(), 'UPI'),
(3, -1000, NOW(), 'UPI'),
(3, -1000, NOW(), 'UPI'),
(3, -1000, NOW(), 'UPI'),
(3, -1000, NOW(), 'UPI'),
(3, -1000, NOW(), 'UPI'),
(3, -1000, NOW(), 'UPI'),
(3, -1000, NOW(), 'UPI'),
(3, -1000, NOW(), 'UPI'),
(3, -1000, NOW(), 'UPI'),
-- Lupin's Transactions
(5, 100000, NOW(), 'NEFT'),
(5, 100000, NOW(), 'NEFT'),
(5, 100000, NOW(), 'NEFT'),
(5, 100000, NOW(), 'NEFT'),
(5, 100000, NOW(), 'NEFT'),
(5, 100000, NOW(), 'NEFT'),
(5, 100000, NOW(), 'NEFT'),
(5, -200000, NOW(), 'NEFT'),
-- Hagrid's Transactions
(6, 10000, NOW(), 'NEFT'),
(6, -2000, NOW(), 'UPI'),
(6, -100, NOW(), 'UPI');
