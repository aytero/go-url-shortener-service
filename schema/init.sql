CREATE TABLE IF NOT EXISTS urls
(
    id SERIAL PRIMARY KEY,
    short_url VARCHAR(10) NOT NULL unique,
    full_url VARCHAR(1024) NOT NULL unique
);