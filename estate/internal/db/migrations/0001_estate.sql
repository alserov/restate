-- +goose Up
CREATE TABLE IF NOT EXISTS estate
(
    id          text PRIMARY KEY,
    title       text,
    description text,
    price       pg_catalog.float8,
    country     text,
    city        text,
    street      text,
    mainImage   text,
    square      pg_catalog.float8,
    floor       int
);

CREATE TABLE IF NOT EXISTS images
(
    estate_id text,
    image     text
);