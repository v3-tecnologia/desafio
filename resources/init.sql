CREATE TABLE public.device (
                               id serial NOT NULL,
                               mac_address macaddr NOT NULL,
                               created_at timestamp NOT NULL,
                               CONSTRAINT device_unique UNIQUE (mac_address),
                               CONSTRAINT device_pk PRIMARY KEY (id)
);


CREATE TABLE public.gps (
                            id serial NOT NULL,
                            device_id int NOT NULL,
                            coordinates public.geography(point, 4326) NOT NULL,
                            "timestamp" timestamp NOT NULL,
                            CONSTRAINT gps_pk PRIMARY KEY (id),
                            CONSTRAINT gps_device_fk FOREIGN KEY (device_id) REFERENCES public.device(id)
);


CREATE TABLE public.gyroscope (
                                  id serial NOT NULL,
                                  device_id int NOT NULL,
                                  x_axis float8 NULL,
                                  y_axis float8 NULL,
                                  z_axis float8 NULL,
                                  "timestamp" timestamp NOT NULL,
                                  CONSTRAINT gyroscope_pk PRIMARY KEY (id),
                                  CONSTRAINT gyroscope_device_fk FOREIGN KEY (device_id) REFERENCES public.device(id)
);

CREATE TABLE public.photo (
                              id serial NOT NULL,
                              device_id int NOT NULL,
                              "path" varchar NOT NULL,
                              "name" varchar NOT NULL,
                              "timestamp" timestamp NOT NULL,
                              CONSTRAINT photo_pk PRIMARY KEY (id),
                              CONSTRAINT photo_device_fk FOREIGN KEY (device_id) REFERENCES public.device(id)
);