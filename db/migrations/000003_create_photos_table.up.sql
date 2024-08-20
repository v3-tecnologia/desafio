CREATE TABLE IF NOT EXISTS photos
(
    id varchar PRIMARY KEY NOT NULL, 
    image_url varchar,
    created_at timestamp DEFAULT now() NOT NULL
);