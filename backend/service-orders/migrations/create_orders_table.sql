CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    client_id INT NOT NULL,
    product_ids JSONB NOT NULL,
    total NUMERIC(10, 2) NOT NULL
);
