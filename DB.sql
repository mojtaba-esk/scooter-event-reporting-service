-- DROP TABLE public.scooters;

CREATE TABLE IF NOT EXISTS public.scooters
(
    uuid character varying(36) COLLATE pg_catalog."default" NOT NULL,
    occupied boolean NOT NULL,
    lat double precision NOT NULL,
    lon double precision NOT NULL,
    last_update timestamp without time zone NOT NULL,
    CONSTRAINT scooters_pkey PRIMARY KEY (uuid)
)
TABLESPACE pg_default;



-- DROP INDEX public.lat;

CREATE INDEX lat
    ON public.scooters USING btree
    (lat ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: lon

-- DROP INDEX public.lon;

CREATE INDEX lon
    ON public.scooters USING btree
    (lon ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: occupied

-- DROP INDEX public.occupied;

CREATE INDEX occupied
    ON public.scooters USING btree
    (occupied ASC NULLS LAST)
    TABLESPACE pg_default;


------------------------------------------

-- DROP TABLE public.trips;

CREATE TABLE IF NOT EXISTS public.trips
(
    uuid character varying(36) COLLATE pg_catalog."default" NOT NULL,
    scooter_uuid character varying(36) COLLATE pg_catalog."default" NOT NULL,
    user_uuid character varying(36) COLLATE pg_catalog."default" NOT NULL,
    start_lat double precision NOT NULL,
    start_lon double precision NOT NULL,
    start_time timestamp without time zone NOT NULL,
    end_lat double precision,
    end_lon double precision,
    end_time timestamp without time zone,
    CONSTRAINT trips_pkey PRIMARY KEY (uuid)
)

TABLESPACE pg_default;

-- DROP INDEX public.scooter_uuid;

CREATE INDEX scooter_uuid
    ON public.trips USING btree
    (scooter_uuid COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: user_uuid

-- DROP INDEX public.user_uuid;

CREATE INDEX user_uuid
    ON public.trips USING btree
    (user_uuid COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;

------------------------------------------

-- DROP TABLE public.clients;

CREATE TABLE IF NOT EXISTS public.clients
(
    uuid character varying(36) COLLATE pg_catalog."default" NOT NULL,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT clients_pkey PRIMARY KEY (uuid)
)

TABLESPACE pg_default;

------------------------------------------

--DROP TABLE public.tracking;

CREATE TABLE IF NOT EXISTS public.tracking
(
    row_id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    scooter_uuid character varying(36) COLLATE pg_catalog."default" NOT NULL,
    lat double precision NOT NULL,
    lon double precision NOT NULL,
    "time" timestamp without time zone NOT NULL,
    CONSTRAINT tracking_pkey PRIMARY KEY (row_id)
)

TABLESPACE pg_default;


-- Index: scooter_uuid_indx

-- DROP INDEX public.scooter_uuid_indx;

CREATE INDEX scooter_uuid_indx
    ON public.tracking USING btree
    (scooter_uuid COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;

------------------------------------------