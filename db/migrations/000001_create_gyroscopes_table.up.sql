CREATE TABLE IF NOT EXISTS gyroscopes
(
    id varchar PRIMARY KEY NOT NULL, 
    x_position float,
    y_position float,
    z_position float,
    created_at timestamp DEFAULT now() NOT NULL
);