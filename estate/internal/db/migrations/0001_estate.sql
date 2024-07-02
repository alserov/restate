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
    images      text[],
    main_image   text,
    square      pg_catalog.float8,
    floor       int
);
