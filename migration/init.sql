CREATE DATABASE urlsshortener;

\connect urlsshortener;

CREATE TABLE urls
(
    id        SERIAL PRIMARY KEY,
    original  VARCHAR(255) UNIQUE,
    shortener VARCHAR(10) UNIQUE
);

INSERT INTO urls (original, shortener) VALUES ('TEST','TEST');