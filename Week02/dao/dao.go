package dao

import (
	"database/sql"
	"log"
	"os"
	"sync"
)

var db *sql.DB
var once sync.Once

/**
sql.DB的设计就是用来作为长连接使用的。不要频繁Open，Close。
比较好的做法是，为每个不同的dataStore各建一个DB对象，保持这些对象Open。
如果需要短连接，那么把DB作为参数传入function，而不要在function中Open，Close。
 */

func init() {
	once.Do(func() {
		driverName := os.Getenv("DB_DRIVER_NAME")
		dataSourceName := os.Getenv("DB_DATA_SOURCE_NAME")
		_db, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			log.Fatal(err)
		}
		db = _db
	})
}
