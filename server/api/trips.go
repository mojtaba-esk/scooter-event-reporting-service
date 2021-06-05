package api

import (
	"encoding/json"
	"log"
	"net/http"
	"scootin/database"
	"scootin/global"
	"scootin/location"
	"scootin/tools"
	"time"

	"github.com/google/uuid"
	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

// type ScooterLocation struct {
// 	Lat  float64   `json:"lat" redis:"lat"`
// 	Lon  float64   `json:"lon" redis:"lon"`
// 	Time time.Time `json:"time" redis:"time"`
// }

/*-------------*/

type Trip struct {
	UUID        string            `json:"uuid" redis:"uuid"`
	ScooterUUID string            `json:"scooter_uuid" redis:"scooter_uuid"`
	UserUUID    string            `json:"user_uuid" redis:"user_uuid"`
	Start       location.Location `json:"start"`
	End         location.Location `json:"end"`
}

/*----------------------*/

/**
* This function implements POST /tripStart
* It starts a new trip
* Body Format: {
	"scooter_uuid": "xxxx",
	"user_uuid": "xxxx",
	"start": {
		"lat": xxxx,
		"lon": xxxx
	}
}
*/
func PostTripStart(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	/*----------*/

	body, err := tools.ReadAll(req.Body)
	if err != nil {
		log.Printf("[ERR  ] PostStartTrip: %s", err.Error())
		http.Error(resp, "bad request", http.StatusBadRequest)
		return
	}

	var inputRecord Trip

	err = json.Unmarshal(body, &inputRecord)
	if err != nil {
		log.Printf("[ERR  ] PostStartTrip: %s", err.Error())
		http.Error(resp, "bad request", http.StatusBadRequest)
		return
	}

	/*----------*/

	// Mark the scooter as occupied
	err = UpdateScooterOccupancy(inputRecord.ScooterUUID, true)
	if err != nil {
		log.Printf("\nError in UpdateScooterOccupancy: %v \nScooterUUID: \n%v", err, inputRecord.ScooterUUID)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	/*----------*/

	tripUUID := uuid.New() // The trip id
	row := database.RowType{
		"uuid":         tripUUID,
		"user_uuid":    inputRecord.UserUUID,
		"scooter_uuid": inputRecord.ScooterUUID,
		"start_lat":    inputRecord.Start.Latitude,
		"start_lon":    inputRecord.Start.Longitude,
		"start_time":   time.Now(),
	}

	_, err = global.DB.Insert("trips", row)
	if err != nil {
		log.Printf("\nError in data insertion: %v \nRow: \n%v", err, row)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)

		// Rollback if fails
		UpdateScooterOccupancy(inputRecord.ScooterUUID, false)
		return
	}

	tools.SendJSON(resp, tripUUID)
}

/*-------------*/
/**
* This function implements POST /tripEnd
* It ends the started trip
* Body Format: {
	"user_uuid": "xxxx",
	"scooter_uuid": "xxxx",
	"end": {
		"lat": xxxx,
		"lon": xxxx
	}
}
*/
func PostTripEnd(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	/*----------*/

	body, err := tools.ReadAll(req.Body)
	if err != nil {
		log.Printf("[ERR  ] PostStartTrip: %s", err.Error())
		http.Error(resp, "bad request", http.StatusBadRequest)
		return
	}

	var inputRecord Trip

	err = json.Unmarshal(body, &inputRecord)
	if err != nil {
		log.Printf("[ERR  ] PostStartTrip: %s", err.Error())
		http.Error(resp, "bad request", http.StatusBadRequest)
		return
	}

	/*----------*/

	conditions := database.RowType{
		"user_uuid":    inputRecord.UserUUID,
		"scooter_uuid": inputRecord.ScooterUUID,
		"end_time":     nil,
	}

	row := database.RowType{
		"end_lat":  inputRecord.End.Latitude,
		"end_lon":  inputRecord.End.Longitude,
		"end_time": time.Now(),
	}

	_, err = global.DB.Update("trips", row, conditions)
	if err != nil {
		log.Printf("\nError in data insertion: %v \nRow: \n%v", err, row)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)

		return
	}

	/*----------------*/

	// Update the scooter's final location to be ready for the next ride
	err = UpdateScootersFinalLocation(inputRecord.ScooterUUID, inputRecord.End.Latitude, inputRecord.End.Longitude)
	if err != nil {
		log.Printf("\nError in UpdateScootersFinalLocation: %v \nScooterUUID: \n%v", err, inputRecord.ScooterUUID)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Mark the scooter as free to use
	err = UpdateScooterOccupancy(inputRecord.ScooterUUID, false)
	if err != nil {
		log.Printf("\nError in UpdateScooterOccupancy: %v \nScooterUUID: \n%v", err, inputRecord.ScooterUUID)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	/*----------------*/

}

/*-------------*/
