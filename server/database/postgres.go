package database

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

/*-----------------*/

func NewPostgresDB(psqlconn string) *sql.DB {

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	return db
}

/*-----------------*/

func (db *Database) PostgresClose() {
	db.SQLConn.Close()
}

/*-----------------*/

func (db *Database) PostgresInsert(table string, fields RowType) (ExecResult, error) {

	SQL := fmt.Sprintf(`INSERT INTO "%s" (`, table)

	var params QueryParams
	values := ""
	paramCounter := 1
	for fieldName, value := range fields {
		SQL += fmt.Sprintf(`"%s",`, fieldName)
		values += fmt.Sprintf(`$%d,`, paramCounter)
		paramCounter++
		params = append(params, value)
	}

	SQL = strings.Trim(SQL, ",")
	values = strings.Trim(values, ",")
	SQL += ") VALUES ( " + values + ")"

	return db.PostgresExec(SQL, params)
}

/*-----------------*/

func (db *Database) PostgresUpdate(table string, fields RowType, conditions RowType) (ExecResult, error) {

	SQL := fmt.Sprintf(`UPDATE "%s" SET `, table)

	var params QueryParams
	paramCounter := 1

	for fieldName, value := range fields {
		SQL += fmt.Sprintf(`"%s" = $%d,`, fieldName, paramCounter)
		paramCounter++
		params = append(params, value)
	}

	SQL = strings.Trim(SQL, ",")
	SQL += " WHERE 1 = 1 "

	for fieldName, value := range conditions {
		SQL += fmt.Sprintf(` AND "%s" = $%d `, fieldName, paramCounter)
		paramCounter++
		params = append(params, value)
	}

	return db.PostgresExec(SQL, params)

}

/*-----------------*/

func (db *Database) PostgresDelete(table string, conditions RowType) (ExecResult, error) {

	SQL := fmt.Sprintf(`DELETE FROM "%s" WHERE 1 = 1 `, table)

	var params QueryParams
	paramCounter := 1
	for fieldName, value := range conditions {
		SQL += fmt.Sprintf(` AND "%s" = $%d `, fieldName, paramCounter)
		paramCounter++
		params = append(params, value)
	}

	return db.PostgresExec(SQL, params)

}

/*-----------------*/

func (db *Database) PostgresExec(query string, params QueryParams) (ExecResult, error) {

	res, err := db.SQLConn.Exec(query, params...)
	if err != nil {
		return ExecResult{}, err
	}

	var output ExecResult

	output.RowsAffected, _ = res.RowsAffected()
	output.LastInsertId, _ = res.LastInsertId()

	return output, nil
}

/*-----------------*/

func (db *Database) PostgresLoad(table string, searchOnFields RowType) (QueryResult, error) {

	SQL := fmt.Sprintf(`SELECT * FROM "%s" WHERE 1 = 1 `, table)

	var params QueryParams
	paramCounter := 1
	for fieldName, value := range searchOnFields {
		SQL += fmt.Sprintf(` AND "%s" = $%d `, fieldName, paramCounter)
		paramCounter++
		params = append(params, value)
	}

	// query := fmt.Sprintf("from(bucket:\"%v\") |> range(start:-1000y) |> filter(fn: (r) => r._measurement == \"%v\")", os.Getenv("PostgresDB_BUCKET"), measurement)
	// return db.PostgresQuery(query)
	return db.PostgresQuery(SQL, params)
}

/*-----------------*/

func (db *Database) PostgresQuery(query string, params QueryParams) (QueryResult, error) {

	var output QueryResult

	rows, err := db.SQLConn.Query(query, params...)
	if err != nil {
		return output, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return output, err
	}

	colCounts := len(columns)
	values := make([]interface{}, colCounts)
	scanArgs := make([]interface{}, colCounts)

	for i := range values {
		scanArgs[i] = &values[i]
	}

	rowCount := 0
	for rows.Next() {

		err = rows.Scan(scanArgs...)
		if err != nil {
			return output, err
		}

		output = append(output, make(RowType, colCounts))
		for i, v := range values {
			output[rowCount][columns[i]] = v
		}
		rowCount++
	}

	return output, nil
}

/*-----------------*/
