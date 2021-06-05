package api

import (
	"log"
	"net/http"
	"os"

	routing "github.com/julienschmidt/httprouter"
)

/*-------------------------*/

func setupRouter() *routing.Router {

	var router = routing.New()

	router.GET("/", IndexPage)
	router.GET("/ui/*file_path", UI)

	router.GET("/clients", CheckAPIKey(GetClients))

	router.GET("/scooters", CheckAPIKey(GetScooters))
	router.GET("/scooters/:uuid", CheckAPIKey(GetScooter))

	router.GET("/scooters/:uuid/location", CheckAPIKey(GetScooterLastKnownLocation))
	router.POST("/scooters/:uuid/location", CheckAPIKey(PostScooterLocation))

	router.POST("/search/freeScooters", CheckAPIKey(PostSearchFreeScooters))
	router.GET("/search/movingScooters", GetMovingScooters) // used for demo only, so no API key needed

	router.POST("/tripStart", CheckAPIKey(PostTripStart))
	router.POST("/tripEnd", CheckAPIKey(PostTripEnd))

	router.GET("/serverReady", GetServerReady)

	return router
}

/*-------------------------*/

// ListenAndServeHTTP serves the APIs and the ui
func ListenAndServeHTTP() {

	router := setupRouter()

	addr := os.Getenv("SERVING_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	log.Printf("[INFO ] Serving on %s", addr)

	log.Fatal(http.ListenAndServe(addr, router))
}

/*-------------------------*/
