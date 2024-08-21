CREATE TABLE IF NOT EXISTS photos
(
    id varchar PRIMARY KEY NOT NULL, 
    image_url varchar,
    is_recognized boolean,
    created_at timestamp DEFAULT now() NOT NULL
);