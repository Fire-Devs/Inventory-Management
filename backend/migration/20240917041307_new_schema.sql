-- +goose Up
CREATE TABLE categories
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL UNIQUE,
    description TEXT
);

CREATE TABLE suppliers
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(255) NOT NULL,
    contact_info TEXT,
    created_at   TIMESTAMPTZ DEFAULT NOW()
);


CREATE TABLE inventory
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    price       NUMERIC(10, 2),
    stock       INT          NOT NULL,
    cover_image TEXT,
    category_id INT REFERENCES categories (id),
    supplier_id INT REFERENCES suppliers (id),
    created_at  TIMESTAMPTZ DEFAULT NOW(),
    updated_at  TIMESTAMPTZ DEFAULT NOW()
);


-- +goose Down
DROP TABLE inventory;
DROP TABLE categories;
DROP TABLE suppliers;
