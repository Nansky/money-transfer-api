CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    payment_type VARCHAR(15) NOT NULL,
    reference_num VARCHAR(20) UNIQUE NOT NULL,
    source_account VARCHAR(20) NOT NULL,
    destination_account VARCHAR(20) NOT NULL,
    gross_amount NUMERIC(15) NOT NULL,
    order_status_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (order_status_id) REFERENCES order_status(id)
);