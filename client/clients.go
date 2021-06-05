package main

import (
	"client/location"
	"client/tools"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

/*-------------*/

func StartClient(client Client) {

	// initial random location for where the client might be
	startLocation := location.Location{51.03879021785863, 13.76123931989416}

	for {

		// Generate two random locations to have the search area
		randomAngle := tools.RandomNumberF(0, 360)
		randomDistance := tools.RandomNumberF(10, 1000)
		newLocation1 := startLocation.Offset(randomDistance, randomAngle)

		randomAngle = tools.RandomNumberF(0, 360)
		randomDistance = tools.RandomNumberF(10, 1000)
		newLocation2 := startLocation.Offset(randomDistance, randomAngle)

		var areaQuery SearchArea
		areaQuery.Start = newLocation1
		areaQuery.End = newLocation2

		freeScooters, _ := SearchForFreeScooters(areaQuery)

		if freeScooters == nil || len(freeScooters) == 0 {
			fmt.Printf("\n\nSorry there is no available scooters at the moment in the selected area!")
			// Let's wait for scooters to get free
			time.Sleep(3 * time.Second)
			continue
		}

		// Just pick a random scooter
		randomIndex := tools.RandomNumberI(0, int64(len(freeScooters))-1)
		selectedScooter := freeScooters[randomIndex]

		trip := Trip{
			Start: location.Location{
				Latitude:  selectedScooter.Lat,
				Longitude: selectedScooter.Lon,
			},
			ScooterUUID: selectedScooter.UUID,
			UserUUID:    client.UUID,
		}

		err := StartTrip(trip)
		if err != nil {
			log.Printf("\n[ERR  ] StartTrip: %v", err)
			return
		}

		const updateInterval = 1 //Seconds
		newScooterLocation := trip.Start
		randomAngle = tools.RandomNumberF(0, 360)
		tripDuration := tools.RandomNumberI(10, 15)
		for i := int64(0); i < tripDuration; {

			time.Sleep(updateInterval * time.Second)
			i += updateInterval

			randomAngle = tools.RandomNumberF(randomAngle-1, randomAngle+1)
			randomDistance := tools.RandomNumberF(30, 70)
			newScooterLocation = newScooterLocation.Offset(randomDistance, randomAngle)

			err := UpdateScootersLocation(selectedScooter.UUID, newScooterLocation)
			if err != nil {
				log.Printf("\n[ERR  ] UpdateScootersLocation: %v", err)
				// return
			}

		}

		time.Sleep(1 * time.Second)

		err = EndTrip(Trip{
			End:         newScooterLocation,
			ScooterUUID: selectedScooter.UUID,
			UserUUID:    client.UUID,
		})
		if err != nil {
			log.Printf("\n[ERR  ] EndTrip: %v", err)
			return
		}

		// Wait before starting the next trip
		time.Sleep(time.Duration(tools.RandomNumberI(2, 5)) * time.Second)
	}
}

/*-------------*/

func GetRandomClients(num int) ([]Client, error) {

	body, err := tools.GetRequest(ENV.SCOOTIN_API_PATH+`clients`, ENV.STATIC_API_KEY)
	if err != nil {
		log.Printf("[ERR ] API call failed: %v", err)
		return []Client{}, err
	}

	var allClients struct {
		Rows []Client `json:"rows"`
	}
	err = json.Unmarshal(body, &allClients)
	if err != nil {
		log.Printf("[ERR  ] Unmarshal: %s", err.Error())
		return []Client{}, err
	}

	output := []Client{}
	totalClients := len(allClients.Rows)

	visitedIndices := map[int64]bool{}
	for i := 0; i < num; {
		randomIndex := tools.RandomNumberI(0, int64(totalClients)-1)
		if _, ok := visitedIndices[randomIndex]; ok {
			continue
		}
		output = append(output, allClients.Rows[randomIndex])
		visitedIndices[randomIndex] = true
		i++
	}

	return output, nil
}

/*-------------*/

func SearchForFreeScooters(query SearchArea) ([]Scooter, error) {

	postBody, err := json.Marshal(query)
	if err != nil {
		log.Printf("[ERR ] Bad input: %v", err)
		return []Scooter{}, err
	}

	body, err := tools.PostRequest(ENV.SCOOTIN_API_PATH+`search/freeScooters`, postBody, ENV.STATIC_API_KEY)
	if err != nil {
		log.Printf("[ERR ] API call failed: %v", err)
		return []Scooter{}, err
	}

	output := []Scooter{}
	err = json.Unmarshal(body, &output)
	if err != nil {
		log.Printf("[ERR  ] Unmarshal: %s", err.Error())
		return []Scooter{}, err
	}

	return output, nil
}

/*-------------*/

func StartTrip(trip Trip) error {

	postBody, err := json.Marshal(trip)
	if err != nil {
		log.Printf("[ERR ] Bad input: %v", err)
		return err
	}

	body, err := tools.PostRequest(ENV.SCOOTIN_API_PATH+`tripStart`, postBody, ENV.STATIC_API_KEY)
	if err != nil {
		log.Printf("[ERR ] API call failed: %v", err)
		return err
	}

	var tripUUID string
	err = json.Unmarshal(body, &tripUUID)
	if err != nil {
		log.Printf("[ERR  ] Unmarshal: %s", err.Error())
		return err
	}

	fmt.Printf("\nTrip started: %v", tripUUID)

	return nil
}

/*-------------*/

func EndTrip(trip Trip) error {

	postBody, err := json.Marshal(trip)
	if err != nil {
		log.Printf("[ERR ] Bad input: %v", err)
		return err
	}

	_, err = tools.PostRequest(ENV.SCOOTIN_API_PATH+`tripEnd`, postBody, ENV.STATIC_API_KEY)
	if err != nil {
		log.Printf("[ERR ] API call failed: %v", err)
		return err
	}

	return nil
}

/*-------------*/

func UpdateScootersLocation(ScooterUUID string, newLocation location.Location) error {

	postBody, err := json.Marshal(newLocation)
	if err != nil {
		log.Printf("[ERR ] Bad input: %v", err)
		return err
	}

	apiPath := fmt.Sprintf("%sscooters/%s/location", ENV.SCOOTIN_API_PATH, ScooterUUID)
	_, err = tools.PostRequest(apiPath, postBody, ENV.STATIC_API_KEY)
	if err != nil {
		log.Printf("[ERR ] API call failed: %v", err)
		return err
	}

	return nil
}

/*-------------*/

func IsServerReady() bool {

	body, err := tools.GetRequest(ENV.SCOOTIN_API_PATH+`serverReady`, "")
	if err != nil {
		return false
	}

	var ready bool
	err = json.Unmarshal(body, &ready)
	if err != nil {
		log.Printf("[ERR  ] Unmarshal: %s", err.Error())
		return false
	}

	return ready
}

/*-------------*/
