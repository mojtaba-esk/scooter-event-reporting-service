package dbinit

import (
	"log"
	"scootin/database"
	"scootin/global"
	"scootin/location"
	"strings"
)

/**
* This function initializes the database
* It checks if DB is not ready and then creates tabels, indices, ...
* then it fills it up with the initial random data
 */
func DatabaseInit() {

	if !NeedToInitDB() {
		return
	}

	/*--------------*/
	log.Printf("Database initialization started.")
	log.Printf("\tCreating Tables and Indices...")

	err := CreateTables()
	if err != nil {
		panic(err)
	}
	log.Printf("Done")

	log.Printf("\tGenerating Random Clients...")
	GenerateRandomClients()
	log.Printf("Done")

	log.Printf("\tGenerating Random Scooters...")
	startLocation := location.Location{51.03879021785863, 13.76123931989416}
	GenerateRandomScooters(500, startLocation)
	log.Printf("Done")

	log.Printf("Database initialization Done.\n\n")
}

/*--------------------------------*/

func NeedToInitDB() bool {

	SQL := `SELECT * FROM "scooters" LIMIT 1;`
	_, err := global.DB.Query(SQL, database.QueryParams{})
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") {
			return true
		}
		panic(err)
	}
	return false
}

/*--------------------------------*/

func CreateTables() error {
	SQList := []string{

		`CREATE TABLE IF NOT EXISTS public.scooters
		(
			uuid character varying(36) COLLATE pg_catalog."default" NOT NULL,
			occupied boolean NOT NULL,
			lat double precision NOT NULL,
			lon double precision NOT NULL,
			last_update timestamp without time zone NOT NULL,
			CONSTRAINT scooters_pkey PRIMARY KEY (uuid)
		)
		TABLESPACE pg_default`,

		`CREATE INDEX lat
			ON public.scooters USING btree
			(lat ASC NULLS LAST)
			TABLESPACE pg_default`,

		`CREATE INDEX lon
			ON public.scooters USING btree
			(lon ASC NULLS LAST)
			TABLESPACE pg_default`,

		`CREATE INDEX occupied
			ON public.scooters USING btree
			(occupied ASC NULLS LAST)
			TABLESPACE pg_default`,

		`CREATE TABLE IF NOT EXISTS public.trips
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

		TABLESPACE pg_default`,
		`CREATE INDEX scooter_uuid
			ON public.trips USING btree
			(scooter_uuid COLLATE pg_catalog."default" ASC NULLS LAST)
			TABLESPACE pg_default`,

		`CREATE INDEX user_uuid
			ON public.trips USING btree
			(user_uuid COLLATE pg_catalog."default" ASC NULLS LAST)
			TABLESPACE pg_default`,

		`CREATE TABLE IF NOT EXISTS public.clients
		(
			uuid character varying(36) COLLATE pg_catalog."default" NOT NULL,
			name character varying(255) COLLATE pg_catalog."default" NOT NULL,
			CONSTRAINT clients_pkey PRIMARY KEY (uuid)
		)
		TABLESPACE pg_default`,

		`CREATE TABLE IF NOT EXISTS public.tracking
		(
			row_id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
			scooter_uuid character varying(36) COLLATE pg_catalog."default" NOT NULL,
			lat double precision NOT NULL,
			lon double precision NOT NULL,
			"time" timestamp without time zone NOT NULL,
			CONSTRAINT tracking_pkey PRIMARY KEY (row_id)
		)

		TABLESPACE pg_default`,

		`CREATE INDEX scooter_uuid_indx
			ON public.tracking USING btree
			(scooter_uuid COLLATE pg_catalog."default" ASC NULLS LAST)
			TABLESPACE pg_default`,
	}

	for _, SQL := range SQList {
		_, err := global.DB.Exec(SQL, database.QueryParams{})
		if err != nil {
			return err
		}
	}

	return nil
}

/*--------------------------------*/
