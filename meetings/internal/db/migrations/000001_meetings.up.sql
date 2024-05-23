CREATE TABLE IF NOT EXISTS meetings
(
    id           text PRIMARY KEY,
    timestamp    timestamp,
    estateID     text,
    visitorPhone text
);