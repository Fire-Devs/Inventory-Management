-- +goose Up
CREATE TABLE ROLES (
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    permissions TEXT[] NOT NULL,
    created_at   TIMESTAMPTZ DEFAULT NOW()
);


-- +goose Down
DROP TABLE ROLES;
