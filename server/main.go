package main

import (
	"fmt"
	"scootin/api"
	"scootin/database"
	"scootin/dbinit"
	"scootin/global"
)

/*--------------------------------*/

func main() {

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		global.ENV.POSTGRES_HOST,
		global.ENV.POSTGRES_PORT,
		global.ENV.POSTGRES_USER,
		global.ENV.POSTGRES_PASSWORD,
		global.ENV.POSTGRES_DB,
	)

	global.DB = database.New(database.Postgres, psqlconn)
	defer global.DB.Close()

	/*--------------*/

	// redisConn := fmt.Sprintf("redis://%s:%s@%s:%s/%s",
	// 	global.ENV.REDIS_USER,
	// 	global.ENV.REDIS_PASSWORD,
	// 	global.ENV.REDIS_HOST,
	// 	global.ENV.REDIS_PORT,
	// 	global.ENV.REDIS_DB,
	// )

	// global.Redis = database.New(database.Redis, redisConn)
	// defer global.Redis.Close()

	/*--------------*/

	dbinit.DatabaseInit()

	global.ServerIsReady = true

	/*--------------*/

	api.ListenAndServeHTTP()
}

/*--------------------------------*/
