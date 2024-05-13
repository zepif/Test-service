-- +migrate Up
CREATE TABLE urlstorage (
    ShortURL TEXT PRIMARY KEY,
    FullURL TEXT NOT NULL
);

-- +migrate Down
DROP TABLE urlstorage;
