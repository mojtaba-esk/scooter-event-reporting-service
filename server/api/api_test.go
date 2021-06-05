package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"scootin/database"
	"scootin/dbinit"
	"scootin/global"
	"strings"
	"testing"
)

/**----------------*/

func TestMain(m *testing.M) {

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		global.ENV.POSTGRES_HOST,
		global.ENV.POSTGRES_PORT,
		global.ENV.POSTGRES_USER,
		global.ENV.POSTGRES_PASSWORD,
		global.ENV.POSTGRES_DB,
	)

	global.DB = database.New(database.Postgres, psqlconn)
	dbinit.DatabaseInit()
	defer global.DB.Close()

	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

/**----------------*/

func TestRouting(t *testing.T) {

	scooterUUID := GetRandomScooter()

	// Test table
	tt := []struct {
		name     string
		api      string
		apiKey   string
		method   string
		postData string
		needle   string // a substring that should be found in the output content
		status   int
		err      string
	}{
		{
			name:   "API Path: /serverReady",
			api:    "/serverReady",
			method: "GET",
			status: 200,
		},
		{
			name:   "API Path: /",
			api:    "/",
			method: "GET",
			status: 200,
		},
		{
			name:   "API Path: /ui/",
			api:    "/ui/",
			method: "GET",
			needle: "<html",
			status: 200,
		},
		{
			name:   "API Path: /clients No API Key",
			api:    "/clients",
			method: "GET",
			needle: "",
			status: 401,
		},
		{
			name:   "API Path: /clients Wrong API key",
			api:    "/clients",
			apiKey: "Wrong API Key",
			method: "GET",
			needle: "",
			status: 401,
		},
		{
			name:   "API Path: /clients",
			api:    "/clients",
			apiKey: global.ENV.STATIC_API_KEY,
			method: "GET",
			needle: "",
			status: 200,
		},
		{
			name:   "API Path: /scooters",
			api:    "/scooters",
			apiKey: global.ENV.STATIC_API_KEY,
			method: "GET",
			needle: "",
			status: 200,
		},
		{
			name:   "API Path: /scooters/notExistingUUID",
			api:    "/scooters/notExistingUUID",
			apiKey: global.ENV.STATIC_API_KEY,
			method: "GET",
			needle: "",
			status: 404,
		},
		{
			name:   "API Path: /scooters/:uuid",
			api:    "/scooters/" + scooterUUID,
			apiKey: global.ENV.STATIC_API_KEY,
			method: "GET",
			needle: scooterUUID,
			status: 200,
		},
		{
			name:   "API Path Post: /scooters/:uuid/location",
			api:    "/scooters/" + scooterUUID + "/location",
			apiKey: global.ENV.STATIC_API_KEY,
			method: "POST",
			postData: `{
				"lat": 51,
				"lon": 32
			}`,
			status: 200,
		},
		{
			name:   "API Path: /scooters/:uuid/location",
			api:    "/scooters/" + scooterUUID + "/location",
			apiKey: global.ENV.STATIC_API_KEY,
			method: "GET",
			needle: `"lat":`,
			status: 200,
		},
		{
			name:   "API Path: /search/freeScooters",
			api:    "/search/freeScooters",
			apiKey: global.ENV.STATIC_API_KEY,
			method: "POST",
			postData: `{
				"start": {
					"lat": 0,
					"lon": 0
				},
				"end": {
					"lat": 500,
					"lon": 500
				}
			}`,
			status: 200,
		},
		{
			name:   "API Path: /tripStart",
			api:    "/tripStart",
			apiKey: global.ENV.STATIC_API_KEY,
			method: "POST",
			postData: `{
				"scooter_uuid": "` + scooterUUID + `",
				"user_uuid": "xxxx",
				"start": {
					"lat": 50,
					"lon": 30
				}
			}`,
			status: 200,
		},
		{
			name:   "API Path: /tripEnd",
			api:    "/tripEnd",
			apiKey: global.ENV.STATIC_API_KEY,
			method: "POST",
			postData: `{
				"scooter_uuid": "` + scooterUUID + `",
				"user_uuid": "xxxx",
				"end": {
					"lat": 80,
					"lon": 19
				}
			}`,
			status: 200,
		},
	}

	/*---------------------*/

	// This Env is set in the Test mode in Dockerfile
	if os.Getenv("EXEC_PATH") == "" {
		t.Fatalf("Env variable `EXEC_PATH` is not set!\n It is required for routing test.")
	}

	/*---------------------*/

	router := setupRouter()

	server := httptest.NewServer(router)
	defer server.Close()

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {

			apiPath := server.URL + tc.api

			req, err := http.NewRequest(tc.method, apiPath, bytes.NewReader([]byte(tc.postData)))
			if err != nil {
				t.Fatalf("Could not sed `%s` request: %v", tc.method, err)
			}

			req.Header.Set("Content-Type", "application/json; charset=UTF-8")

			if tc.apiKey != "" {
				req.Header.Set("X-API-KEY", tc.apiKey)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("Could not receive any response: %v", err)
			}
			defer res.Body.Close()

			content, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("Could not read the content: %v", err)
			}

			if res.StatusCode != tc.status {
				t.Fatalf("Request failed: `%v` ==> %v", apiPath, res.Status)
			}

			if tc.err != "" && tc.err != string(content) {
				t.Errorf("Expected error message %q; got %q", tc.err, string(content))
			}

			if tc.needle != "" && !strings.Contains(string(content), tc.needle) {
				t.Fatalf("Expected to find %q in the content", tc.needle)
			}
		})
	}
}

/*----------------*/

func GetRandomScooter() string {

	/*------*/

	SQL := `SELECT "uuid" FROM "scooters" ORDER BY RANDOM() LIMIT 1`

	rows, err := global.DB.Query(SQL, database.QueryParams{})
	if err != nil {
		log.Printf("Error in db query: %v", err)
		return ""
	}

	if rows == nil || len(rows) == 0 {
		return ""
	}

	return rows[0]["uuid"].(string)
}

/*----------------*/
