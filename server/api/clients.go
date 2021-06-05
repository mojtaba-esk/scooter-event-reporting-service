package api

import (
	"log"
	"math"
	"net/http"
	"scootin/database"
	"scootin/global"
	"scootin/tools"

	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

type Client struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

/*-------------*/

/*
* This function implements GET /clients
 */
func GetClients(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	limitOffset := tools.GetLimitOffset(req)

	/*------*/

	totalRows := int64(0)
	{
		SQL := `SELECT COUNT(*) AS "total" FROM "clients"`
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
				"clients"
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
