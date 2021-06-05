package api

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"scootin/database"
	"scootin/global"
	"scootin/location"
	"scootin/tools"

	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

type SearchArea struct {
	Start location.Location `json:"start" redis:"start"`
	End   location.Location `json:"end" redis:"end"`
}

/*----------------------*/

/**
* This function implements POST /search/freeScooters
* Query: {
	"start": {
		"lat": xxxxx,
		"lon": xxxxx,
	},
	"end": {
		"lat": xxxxx,
		"lon": xxxxx,
	}
}
*
*/
func PostSearchFreeScooters(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	/*----------*/

	body, err := tools.ReadAll(req.Body)
	if err != nil {
		log.Printf("[ERR  ] GetSearchFreeScooters: %s", err.Error())
		http.Error(resp, "bad request", http.StatusBadRequest)
		return
	}

	var query SearchArea

	err = json.Unmarshal(body, &query)
	if err != nil {
		log.Printf("[ERR  ] GetSearchFreeScooters: %s", err.Error())
		http.Error(resp, "bad request", http.StatusBadRequest)
		return
	}

	/*----------*/

	var minLat, maxLat, minLon, maxLon float64

	minLat = math.Min(query.Start.Latitude, query.End.Latitude)
	maxLat = math.Max(query.Start.Latitude, query.End.Latitude)

	minLon = math.Min(query.Start.Longitude, query.End.Longitude)
	maxLon = math.Max(query.Start.Longitude, query.End.Longitude)

	/*----------*/

	SQL := `SELECT
				"uuid", "lat", "lon"
			FROM 
				"scooters"
			WHERE
				"occupied" = false			AND
				"lat" BETWEEN $1 AND $2		AND
				"lon" BETWEEN $3 AND $4
			LIMIT $5`

	rows, err := global.DB.Query(SQL, database.QueryParams{minLat, maxLat, minLon, maxLon, global.RowsPerPage})
	if err != nil {
		log.Printf("Error in db query: %v\n SQL: %v", err, SQL)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rows == nil || len(rows) == 0 {
		http.Error(resp, "no available scooters found!", http.StatusNotFound)
		return
	}

	tools.SendJSON(resp, rows)
}

/*-------------*/
/**
* This function implements GET /search/movingScooters
* and it is mainly used for visualization
 */

func GetMovingScooters(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	//TODO: Get it from Redis in future

	SQL := `SELECT 
				"s"."uuid", "t"."lat", "t"."lon", "t"."time" 
			FROM 
				"scooters" AS "s",
				(
					SELECT scooter_uuid, MAX(row_id) AS last_row 
					FROM "tracking" 
					GROUP BY scooter_uuid
					) AS l
				INNER JOIN tracking AS t
				ON l.last_row = t.row_id 
			WHERE
				s.occupied = true	AND
				s.uuid = t.scooter_uuid
			LIMIT $1`

	rows, err := global.DB.Query(SQL, database.QueryParams{global.RowsPerPage})
	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rows == nil || len(rows) == 0 {
		http.Error(resp, "No moving scooters found!", http.StatusNotFound)
		return
	}

	tools.SendJSON(resp, rows)
}

/*-------------*/
