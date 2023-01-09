package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "practice_db"
)

func connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbcon, err := sql.Open("postgres", psqlInfo)

	if err != nil || dbcon.Ping() != nil {
		panic(err.Error())
	}

	defer dbcon.Close()

	return dbcon
}

func GetMessage(id uint64) {

}

func SaveMessage() {

}

func UpdateMessage() {

}

func DeleteMessage() {

}
