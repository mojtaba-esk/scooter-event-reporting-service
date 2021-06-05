package api

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"scootin/database"
	"scootin/global"
	"scootin/tools"
	"time"

	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

type Scooter struct {
	UUID       string    `json:"uuid" redis:"uuid"`
	Occupied   bool      `json:"occupied" redis:"occupied"`
	Lat        float64   `json:"lat" redis:"lat"`
	Lon        float64   `json:"lon" redis:"lon"`
	LastUpdate time.Time `json:"last_update" redis:"last_update"`
}

/*----------------------*/

/*
* This function implements GET /scooters
 */
func GetScooters(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	limitOffset := tools.GetLimitOffset(req)

	/*------*/

	totalRows := int64(0)
	{
		SQL := `SELECT COUNT(*) AS "total" FROM "scooters"`
		rows, err := global.DB.Query(SQL, database.QueryParams{})
		if err != nil {
			log.Printf("Error in db query: %v", err)
			http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		totalRows = rows[0]["total"].(int64)
	}

	totalPages := int64(math.Ceil(float64(totalRows) / float64(global.RowsPerPage)))
	pagination := map[string]interface{}{
		"current_page":  limitOffset.Page,
		"total_pages":   totalPages,
		"total_entries": totalRows,
	}

	/*------*/

	SQL := `SELECT *
			FROM 
				"scooters"
			LIMIT $1 OFFSET $2`

	rows, err := global.DB.Query(SQL, database.QueryParams{limitOffset.Limit, limitOffset.Offset})
	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, map[string]interface{}{"pagination": pagination, "rows": rows})
}

/*-------------*/

/*
* This function implements GET /scooters/:uuid
 */
func GetScooter(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	uuid := params.ByName("uuid")

	/*------*/

	SQL := `SELECT * FROM "scooters" WHERE "uuid" = $1`

	rows, err := global.DB.Query(SQL, database.QueryParams{uuid})
	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rows == nil || len(rows) == 0 {
		http.Error(resp, "Scooter not found!", http.StatusNotFound)
		return
	}

	tools.SendJSON(resp, rows[0])
}

/*-------------*/

/**
* This function implements GET /scooters/:uuid/location
 */
func GetScooterLastKnownLocation(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	uuid := params.ByName("uuid")

	/*------*/
	//TODO: Get it from Redis in future

	SQL := `SELECT 
				"lat", "lon", "time"
			FROM 
				"tracking" 
			WHERE 
				"scooter_uuid" = $1 
			ORDER BY "row_id" 
			LIMIT 1`

	rows, err := global.DB.Query(SQL, database.QueryParams{uuid})
	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rows == nil || len(rows) == 0 {
		http.Error(resp, "Location not found!", http.StatusNotFound)
		return
	}

	tools.SendJSON(resp, rows[0])
}

/*-------------*/

/**
* This function implements POST /scooters/:uuid/location
* Body: {
	"lat": xxxxxxx,
	"lon": xxxxxxx
}
*/
func PostScooterLocation(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	uuid := params.ByName("uuid")

	/*----------*/

	body, err := tools.ReadAll(req.Body)
	if err != nil {
		log.Printf("[ERR  ] PostScooterLocation: %s", err.Error())
		http.Error(resp, "bad request", http.StatusBadRequest)
		return
	}

	var inputRecord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	}

	err = json.Unmarshal(body, &inputRecord)
	if err != nil {
		log.Printf("[ERR  ] PostScooterLocation: %s", err.Error())
		http.Error(resp, "bad request", http.StatusBadRequest)
		return
	}

	/*----------*/

	row := database.RowType{
		"scooter_uuid": uuid,
		"lat":          inputRecord.Lat,
		"lon":          inputRecord.Lon,
		"time":         time.Now(),
	}

	_, err = global.DB.Insert("tracking", row)
	if err != nil {
		log.Printf("\nError in data insertion: %v \nRow: \n%v", err, row)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

/*-------------*/

func UpdateScooterOccupancy(uuid string, occupied bool) error {

	condition := database.RowType{
		"uuid": uuid,
	}

	row := database.RowType{
		"occupied":    occupied,
		"last_update": time.Now(),
	}

	_, err := global.DB.Update("scooters", row, condition)
	return err
}

/*-------------*/

func UpdateScootersFinalLocation(uuid string, lat float64, lon float64) error {

	condition := database.RowType{
		"uuid": uuid,
	}

	row := database.RowType{
		"lat":         lat,
		"lon":         lon,
		"last_update": time.Now(),
	}

	_, err := global.DB.Update("scooters", row, condition)
	return err
}

/*-------------*/
