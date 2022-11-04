package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host       = "localhost"
	port       = 5432
	user       = "postgres"
	password   = "1"
	dbname     = "quizbackend"
	sslenabled = "disable" //disable or required
)

var (
	Db *sql.DB
)

func StartDB() {
	var err error
	Db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslenabled))
	if err != nil {
		log.Fatal(err) //temp replace with real logging system
	}
}
