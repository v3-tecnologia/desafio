CREATE TABLE IF NOT EXISTS gps
(
    id varchar PRIMARY KEY NOT NULL, 
    latitude float,
    longitude float,
    created_at timestamp DEFAULT now() NOT NULL
);