CREATE TABLE IF NOT EXISTS OFFER_METADATA (
    ID SERIAL PRIMARY KEY,
    TITLE VARCHAR(255) NOT NULL,
    PRODUCT VARCHAR(255) NOT NULL,
    DESCRIPTION VARCHAR(255) NOT NULL,
);

INSERT INTO OFFER_METADATA(TITLE, PRODUCT, DESCRIPTION) 
VALUES 
('Mutual Funds Management and Demat Account', 'Investment', 'Invest in top-performing mutual funds and manage your Demat account seamlessly.'),
('Gold Credit Card', 'Credit Card', 'Exclusive Gold Credit Card with premium benefits, rewards, and cashback on transactions.'),
('Fixed Deposit', 'Deposits', 'Secure your future with our high-interest Fixed Deposit schemes.'),
('Personal Loan', 'Loans', 'Get an instant personal loan with minimal documentation and attractive interest rates.'),
('RuPay UPI Credit Card', 'Credit Card', 'Enjoy seamless UPI transactions with the RuPay Credit Card, combining convenience and security.'),
('Senior Citizen Fixed Deposit', 'Deposits', 'Special Fixed Deposit scheme for senior citizens with higher interest rates and flexible tenures.'),
('Senior Citizen Recurring Deposit', 'Deposits', 'A safe and convenient Recurring Deposit plan tailored for senior citizens with attractive interest rates.');
