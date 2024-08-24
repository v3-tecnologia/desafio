CREATE TABLE IF NOT EXISTS gyroscopes
(
    id          UUID PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    model       VARCHAR(50)  NOT NULL,
    mac_address VARCHAR(17)  NOT NULL,
    x           DOUBLE PRECISION         DEFAULT 0,
    y           DOUBLE PRECISION         DEFAULT 0,
    z           DOUBLE PRECISION         DEFAULT 0,
    timestamp   TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS gps
(
    id          UUID PRIMARY KEY,
    latitude    DOUBLE PRECISION NOT NULL,
    longitude   DOUBLE PRECISION NOT NULL,
    mac_address VARCHAR(17)      NOT NULL,
    timestamp   TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);