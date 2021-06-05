package tools

import (
	"math/rand"
	"time"
)

/*----------------*/

func RandomNumberF(rangeLower float64, rangeUpper float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rangeLower + rand.Float64()*(rangeUpper-rangeLower)
}

/*----------------*/

func RandomNumberI(rangeLower int64, rangeUpper int64) int64 {

	rand.Seed(time.Now().UnixNano())
	return rangeLower + rand.Int63n(rangeUpper-rangeLower+1)
}

/*----------------*/
