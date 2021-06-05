package api

import (
	"log"
	"net/http"
	"scootin/global"

	routing "github.com/julienschmidt/httprouter"
)

/*----------------*/

// CheckAPIKey checks if the given request is valid for the API call
func CheckAPIKey(endpoint routing.Handle) routing.Handle {

	return func(resp http.ResponseWriter, req *http.Request, params routing.Params) {

		apiKey := req.Header.Get("X-API-KEY")

		if apiKey == global.ENV.STATIC_API_KEY {

			endpoint(resp, req, params)

		} else {
			log.Printf("[ERR  ] invalid API Key: %q", apiKey)
			http.Error(resp, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
	}
}

/*---------------------*/
