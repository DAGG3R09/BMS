package dao

import (
	"database/sql"
	"log"

	// postgres driver
	_ "github.com/lib/pq"
)

var connStr = "user=dagger dbname=bms sslmode=disable"
var db, err = sql.Open("postgres", connStr)

//InitiateDatabaseConnection initiates database session
func InitiateDatabaseConnection() {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

// CloseDatabaseConnection closes database session
func CloseDatabaseConnection() {
	db.Close()
}
