-- +migrate Up  
CREATE TABLE links (
    id SERIAL PRIMARY KEY,
    full_url TEXT NOT NULL,
    short_url TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +migrate Down
DROP TABLE links;
