-- Create users table
CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    phone_number VARCHAR(15) UNIQUE NOT NULL,
    address TEXT,
    pin TEXT,
    balance INT DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Create transactions table
CREATE TABLE transactions (
    transaction_id UUID PRIMARY KEY,
    top_up_id UUID,
    payment_id UUID,
    transfer_id UUID,
    user_id UUID,
    amount INT,
    balance_before INT,
    balance_after INT,
    transaction_type VARCHAR(50),
    remarks TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Insert into users table
INSERT INTO users (user_id, first_name, last_name, phone_number, address, pin, balance)
VALUES
    ('f8f332eb-79a6-4e2f-b4a1-89e9c6f2f5ea', 'firstname', 'lastname', '082345678900', 'Jakarta', '123678', 0),
    ('37e30b89-7c8b-4b57-939f-aae0d79e5aeb', 'lala', 'nana', '085876543221', 'Bogor', '567678', 0);

-- Insert into transactions table
INSERT INTO transactions (transaction_id, top_up_id, payment_id, transfer_id, user_id, amount, balance_before, balance_after, transaction_type, remarks)
VALUES
    ('53d9e1e4-937e-4b23-bf01-51056e76f0cf', NULL, 'a9a7e2e9-0f0c-4b8d-bd1b-9511e775c58a', NULL, 'f8f332eb-79a6-4e2f-b4a1-89e9c6f2f5ea', 500, 1000, 500, 'DEBIT', 'Payment for services'),
    ('af03d899-837c-4e42-9a4a-1a6e2fe8e3ea', '7e96a591-016b-4465-9d20-0a0f154d84d8', NULL, NULL, '37e30b89-7c8b-4b57-939f-aae0d79e5aeb', 200, 500, 700, 'CREDIT', 'Top-up from credit card');


-- Select all users
SELECT * FROM users;

-- Select all transactions
SELECT * FROM transactions;

-- Select transactions by user_id
SELECT * FROM transactions WHERE user_id = 'f8f332eb-79a6-4e2f-b4a1-89e9c6f2f5ea';

