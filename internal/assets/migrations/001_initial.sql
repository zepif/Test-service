-- +migrate Up
CREATE TABLE urlstorage (
    id SERIAL PRIMARY KEY,
    full_url TEXT NOT NULL,
    short_url TEXT UNIQUE NOT NULL
);

-- +migrate Down
DROP TABLE urlstorage;
