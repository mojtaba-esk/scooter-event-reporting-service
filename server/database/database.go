package database

import (
	"database/sql"
)

type DBType int

const (
	Postgres DBType = iota
	Redis
	// InfluxDB
	// MySQL
	// MongoDB
)

type Database struct {
	Type    DBType // postgres, influx, mysql, sqlite,...
	SQLConn *sql.DB
	// Redis   *redis.Conn
	// MySQLConn ...
}

type ExecResult struct {
	RowsAffected int64
	LastInsertId int64
}

type RowType map[string]interface{}
type QueryResult []RowType
type QueryParams []interface{}

/*-----------------------*/

func New(DatabaseType DBType, params ...string) *Database {
	var newDB Database

	newDB.Type = DatabaseType

	switch DatabaseType {
	case Postgres:
		if len(params) == 0 {
			return nil
		}
		newDB.SQLConn = NewPostgresDB(params[0])

		// case Redis:
		// 	if len(params) == 0 {
		// 		return nil
		// 	}
		// 	newDB.Redis = NewRedis(params[0])
	}

	return &newDB
}

/*-----------------------*/

func (db *Database) Close() {
	switch db.Type {
	case Postgres:
		db.PostgresClose()
		// case Redis:
		// 	db.RedisClose()
	}
}

/*-----------------------*/

func (db *Database) Insert(table string, fields RowType) (ExecResult, error) {

	switch db.Type {
	// case Redis:
	// 	return db.RedisInsert(table /*key*/, fields)
	case Postgres:
		return db.PostgresInsert(table, fields)
	}

	return ExecResult{}, nil //TODO: provide a useful error here
}

/*-----------------------*/

func (db *Database) Update(table string, fields RowType, conditions RowType) (ExecResult, error) {

	switch db.Type {
	case Postgres:
		return db.PostgresUpdate(table, fields, conditions)
	}

	return ExecResult{}, nil //TODO: provide a useful error here
}

/*-----------------------*/

func (db *Database) Delete(table string, conditions RowType) (ExecResult, error) {

	switch db.Type {
	case Postgres:
		return db.PostgresDelete(table, conditions)
	}

	return ExecResult{}, nil //TODO: provide a useful error here
}

/*-----------------------*/

func (db *Database) Load(table string, searchOnFields RowType) (QueryResult, error) {

	switch db.Type {
	// case Redis:
	// 	return db.RedisLoad(table /*Key*/, searchOnFields)
	case Postgres:
		return db.PostgresLoad(table, searchOnFields)
	}

	return QueryResult{}, nil //TODO: provide a useful error here

}

/*-----------------------*/

func (db *Database) Query(query string, params QueryParams) (QueryResult, error) {

	switch db.Type {
	case Postgres:
		return db.PostgresQuery(query, params)
	}

	return QueryResult{}, nil //TODO: provide a useful error here

}

/*-----------------------*/

func (db *Database) Exec(query string, params QueryParams) (ExecResult, error) {

	switch db.Type {
	case Postgres:
		return db.PostgresExec(query, params)
	}

	return ExecResult{}, nil //TODO: provide a useful error here

}

/*-----------------------*/
