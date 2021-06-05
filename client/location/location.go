package location

import "math"

type Location struct {
	Latitude  float64 `json:"lat" redis:"lat"`
	Longitude float64 `json:"lon" redis:"lon"`
}

/*--------------------*/
/**
* This function calculates distance between two points in meter
 */
func (l1 Location) Distance(l2 Location) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, earthRadius float64
	la1 = l1.Latitude * math.Pi / 180
	lo1 = l1.Longitude * math.Pi / 180
	la2 = l2.Latitude * math.Pi / 180
	lo2 = l2.Longitude * math.Pi / 180

	earthRadius = 6378100 // Earth radius in meters

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * earthRadius * math.Asin(math.Sqrt(h))
}

/*--------------------*/

/**
* This function receives a distance and an angle
* and calculates the new coordination with the given distance and angle from the current location
 */
func (l Location) Offset(distance float64, angle float64) Location {

	distanceNorth := math.Sin(angle) * distance
	distanceEast := math.Cos(angle) * distance
	earthRadius := float64(6378100)
	newLat := l.Latitude + (distanceNorth/earthRadius)*180/math.Pi
	newLon := l.Longitude + (distanceEast/(earthRadius*math.Cos(newLat*180/math.Pi)))*180/math.Pi

	return Location{newLat, newLon}
}

/*--------------------*/

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

/*--------------------*/
