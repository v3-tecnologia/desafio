CREATE TABLE IF NOT EXISTS public.gps (
	gpsid serial4 NOT NULL,
	mac varchar(255) NOT NULL,
	latitude varchar(255) NULL,
	longitude varchar(255) NULL,
	datatimestamp int4 NULL,
	created int4 NULL,
	CONSTRAINT gps_pkey PRIMARY KEY (gpsid)
);

CREATE TABLE IF NOT EXISTS public.gyroscope (
	gyroscopeid serial4 NOT NULL,
	mac varchar(255) NOT NULL,
	xcoord float8 NULL,
	ycoord float8 NULL,
	zcoord float8 NULL,
	datatimestamp int4 NULL,
	created int4 NULL,
	CONSTRAINT gyroscope_pkey PRIMARY KEY (gyroscopeid)
);

CREATE TABLE IF NOT EXISTS public.photo (
	photoid serial4 NOT NULL,
	mac varchar(255) NOT NULL,
	photo bytea NULL,
	datatimestamp int4 NULL,
	created int4 NULL,
	CONSTRAINT photo_pkey PRIMARY KEY (photoid)
);