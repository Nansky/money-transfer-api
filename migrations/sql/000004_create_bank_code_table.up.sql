CREATE TABLE IF NOT EXISTS bank_code (
    id SERIAL PRIMARY KEY,
    bank_code VARCHAR(10) NOT NULL,
    bank_name VARCHAR(25) NOT NULL
);

-- seed
INSERT INTO bank_code (bank_code, bank_name) VALUES 
    ('999', 'testbank'),
    ('111', 'testbank2');
