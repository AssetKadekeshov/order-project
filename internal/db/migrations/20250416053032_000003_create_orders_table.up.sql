CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    customer_name TEXT NOT NULL,
    products TEXT[] NOT NULL,
    total_quantity INTEGER NOT NULL,
    status TEXT NOT NULL
);
