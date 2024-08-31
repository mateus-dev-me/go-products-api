CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    category VARCHAR(255),
    description TEXT,
    quantity INT NOT NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS products;


