CREATE TABLE IF NOT EXISTS public.gyroscope_data (
	id bigserial NOT NULL,
	macaddress varchar(255) NOT NULL,
	x_coord float8 NULL,
	y_coord float8 NULL,
	z_coord float8 NULL,
	data_timestamp int4 NULL,
	created int4 NULL,
	CONSTRAINT gyroscope_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.gps_data (
	id bigserial NOT NULL,
	macaddress varchar(255) NOT NULL,
	latitude varchar(255) NULL,
	longitude varchar(255) NULL,
	data_timestamp int4 NULL,
	created int4 NULL,
	CONSTRAINT gps_pkey PRIMARY KEY (id)
);



CREATE TABLE IF NOT EXISTS public.photo (
	id bigserial NOT NULL,
	macaddress varchar(255) NOT NULL,
	photo bytea NULL,
	data_timestamp int4 NULL,
	created int4 NULL,
	CONSTRAINT photo_pkey PRIMARY KEY (id)
);