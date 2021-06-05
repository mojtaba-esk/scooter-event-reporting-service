package api

import (
	"net/http"
	"scootin/global"
	"scootin/tools"

	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

/**
* This function implements GET /serverReady
* This function indicates if the server is ready to be used by client Apps after it is just launched
 */
func GetServerReady(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	tools.SendJSON(resp, global.ServerIsReady)
}

/*-------------*/
