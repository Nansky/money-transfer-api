CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    account_number VARCHAR(20) UNIQUE NOT NULL,
    balance  NUMERIC(15) NOT NULL,
    account_name VARCHAR(50) NOT NULL,
    bank_code varchar(25) NOT NULL
);

CREATE INDEX ON accounts (account_number);

-- seed
INSERT INTO accounts (account_number, balance, account_name, bank_code) VALUES 
    ('034101056895506', 15000000, 'Si Budi', 'testbank'),
    ('066123490', 500000, 'Pak Udin', 'testbank'),
    ('0110632321111111', 600000, 'Cloud Strife', 'testbank'),
    ('0000412421', 140000000, 'Master Test Bank', 'testbank'),
    ('86489542', 950000, 'Bank Sat', 'testbank'),
    ('7148730038', 50000000, 'Habib Jafar', 'testbank'),
    ('3941000000030', 2500000, 'Test Akun', 'testbank'),
    ('111001110963', 780000, 'Nina Erdiani', 'testbank'),
    ('6711111100', 1435000, 'Bu Nada', 'testbank'),
    ('6013010599', 660500, 'Squall Leonhart', 'testbank');