package dbinit

import (
	"log"
	"scootin/database"
	"scootin/global"
	"scootin/location"
	"scootin/tools"
	"time"

	"github.com/google/uuid"
)

/*------------------------*/

func GenerateRandomScooters(numOfScooters int, initLocation location.Location) error {

	for i := 0; i < numOfScooters; i++ {

		randomAngle := tools.RandomNumberF(0, 360)
		randomDistance := tools.RandomNumberF(10, 2000)

		newLocation := initLocation.Offset(randomDistance, randomAngle)

		row := database.RowType{
			"uuid":        uuid.New(),
			"occupied":    false,
			"lat":         newLocation.Latitude,
			"lon":         newLocation.Longitude,
			"last_update": time.Now(),
		}

		_, err := global.DB.Insert("scooters", row)
		if err != nil {
			log.Printf("\nError in data insertion: %v \nRow: \n%v", err, row)
		}
	}

	return nil
}

/*----------------*/
