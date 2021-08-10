CREATE DATABASE urlsshortener;

\connect urlsshortener;

CREATE TABLE urlsshortener.urls
(
    id        SERIAL PRIMARY KEY,
    original  VARCHAR(255) UNIQUE,
    shortener VARCHAR(10) UNIQUE
);

INSERT INTO urlsshortener.urls (original, shortener) VALUES ('TEST','TEST');