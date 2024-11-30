CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE customers_address (
    id SERIAL PRIMARY KEY,
    customer INT REFERENCES customers (id),
    contact_number VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    address_1 TEXT NOT NULL,
    address_2 TEXT,
    city VARCHAR(255) NOT NULL,
    state VARCHAR(255) NOT NULL,
    postal_code VARCHAR(255) NOT NULL,
    isShipping BOOLEAN NOT NULL DEFAULT FALSE,
    isBilling BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    quantity INT NOT NULL DEFAULT 1,
    customer INT REFERENCES customers (id),
    product INT REFERENCES inventory (id),
    shipping_address INT REFERENCES customers_address (id),
    billing_address INT REFERENCES customers_address (id),
    created_at TIMESTAMPTZ DEFAULT NOW()
);
