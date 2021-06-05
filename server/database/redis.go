package database

// import (
// 	"fmt"

// 	"github.com/gomodule/redigo/redis"
// )

// /*-----------------*/

// // Unix Timestamp
// type Timestamp int64

// // type Timestamp time.Time

// // func (t *Timestamp) RedisScan(x interface{}) error {
// // 	bs, ok := x.([]byte)
// // 	if !ok {
// // 		return fmt.Errorf("expected []byte, got %T", x)
// // 	}
// // 	tt, err := time.Parse(time.RFC3339, string(bs))
// // 	if err != nil {
// // 		return err
// // 	}
// // 	*t = Timestamp(tt)
// // 	return nil
// // }

// /*-----------------*/

// func NewRedis(connStr string) *redis.Conn {

// 	db, err := redis.DialURL(connStr)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return &db
// }

// /*-----------------*/

// func (db *Database) RedisClose() {
// 	(*db.Redis).Close()
// }

// /*-----------------*/

// func (db *Database) RedisInsert(key string, fields RowType) (ExecResult, error) {
// 	_, err := db.RedisSave(key, fields)
// 	if err != nil {
// 		return ExecResult{}, err
// 	}
// 	return ExecResult{}, nil
// }

// /*-----------------*/

// func (db *Database) RedisSave(key string, fields interface{}) (ExecResult, error) {

// 	res, err := (*db.Redis).Do("HMSET", redis.Args{key}.AddFlat(fields)...)

// 	if err != nil {
// 		return ExecResult{}, err
// 	}

// 	fmt.Printf("Result: %+v", res)

// 	return ExecResult{}, nil

// }

// // /*-----------------*/

// func (db *Database) RedisLoad(key string, searchOnFields RowType) (QueryResult, error) {
// 	return QueryResult{}, nil
// }

// /*----------------*/

// func (db *Database) RedisGet(key string, resultStruct interface{}) error {

// 	reply, err := redis.Values((*db.Redis).Do("HGETALL", key))
// 	if err != nil {
// 		return err
// 	}

// 	err = redis.ScanStruct(reply, resultStruct)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("reply: %+v", reply)
// 	fmt.Printf("resultStruct: %+v", resultStruct)

// 	return nil

// }

// /*----------------*/

// func (db *Database) RedisGetRange(query interface{}, resultStruct interface{}) error {

// 	reply, err := redis.Values((*db.Redis).Do("ZRANGEBYSCORE", redis.Args{}.AddFlat(query)...))
// 	if err != nil {
// 		fmt.Printf("\n\tError: %+v\n", err)
// 		return err
// 	}

// 	// err = redis.ScanStruct(reply, resultStruct)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	fmt.Printf("reply: %+v", reply)

// 	return nil

// }

// /*-----------------*/
