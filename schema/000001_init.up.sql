CREATE TABLE urls
(
    id        SERIAL PRIMARY KEY,
    original  VARCHAR(255) UNIQUE,
    shortener VARCHAR(10) UNIQUE
);
