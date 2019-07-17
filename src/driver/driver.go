package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// global variable available in the whole package
var db *sql.DB

// database connection info:
const (
	host   = "localhost"
	port   = 5432
	user   = "jonathancannon"
	dbname = "test"
)

func SQLConnect() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	db, _ = sql.Open("postgres", psqlInfo)

	err := db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Established connection to database: ", dbname)

	return db
}
