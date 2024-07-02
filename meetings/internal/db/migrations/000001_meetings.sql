CREATE TABLE IF NOT EXISTS meetings
(
    id           text PRIMARY KEY,
    timestamp    timestamp,
    estate_id     text,
    visitor_phone text
);