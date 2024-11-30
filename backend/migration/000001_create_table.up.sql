CREATE TABLE inventory
(
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(255) NOT NULL,
    stock         INT          NOT NULL,
    description   TEXT,
    cover_image   TEXT,
    created_at    TIMESTAMPTZ DEFAULT NOW(),
    updated_at    TIMESTAMPTZ DEFAULT NOW()
);
CREATE TABLE categories
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL UNIQUE,
    inventory   INT REFERENCES inventory (id),
    description TEXT
);

CREATE TABLE suppliers
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(255) NOT NULL,
    inventory   INT REFERENCES inventory (id),
    contact_info TEXT,
    created_at   TIMESTAMPTZ DEFAULT NOW()
);


CREATE TABLE price
(
    id         SERIAL PRIMARY KEY,
    inventory  INT REFERENCES inventory (id),
    currency   VARCHAR(3)     NOT NULL,
    amount     NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);









