CREATE TABLE IF NOT EXISTS photos
(
    id varchar PRIMARY KEY NOT NULL, 
    image_url varchar,
    is_recognized boolean,
    amount_of_faces_detected integer,
    confidence_mean float,
    created_at timestamp DEFAULT now() NOT NULL
);