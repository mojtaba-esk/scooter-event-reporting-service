package global

import (
	"os"
	"scootin/database"
)

/*-------------*/

var DB *database.Database    // initiated in the main package
var Redis *database.Database // initiated in the main package

const RowsPerPage = 200 // This is the number of rows that APIs show per page

/*-------------*/

var ServerIsReady bool

/*-------------*/

var ENV struct {
	SERVING_ADDR   string
	STATIC_API_KEY string

	POSTGRES_DB       string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_PORT     string
	POSTGRES_HOST     string

	REDIS_USER     string
	REDIS_PASSWORD string
	REDIS_HOST     string
	REDIS_PORT     string
	REDIS_DB       string
}

/*-------------*/

func init() {

	/*----------*/

	ENV.SERVING_ADDR = os.Getenv("SERVING_ADDR")
	ENV.STATIC_API_KEY = os.Getenv("STATIC_API_KEY")

	ENV.POSTGRES_DB = os.Getenv("POSTGRES_DB")
	ENV.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	ENV.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	ENV.POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	ENV.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")

	ENV.REDIS_USER = os.Getenv("REDIS_USER")
	ENV.REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	ENV.REDIS_HOST = os.Getenv("REDIS_HOST")
	ENV.REDIS_PORT = os.Getenv("REDIS_PORT")
	ENV.REDIS_DB = os.Getenv("REDIS_DB")

	/*----------*/

	ServerIsReady = false
}

/*-------------*/
