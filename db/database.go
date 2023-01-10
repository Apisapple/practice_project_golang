package db

import (
	"database/sql"
	"fmt"

	model "example.com/practice/model"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbcon, err := sql.Open("postgres", psqlInfo)

	if err != nil || dbcon.Ping() != nil {
		panic(err.Error())
	}

	return dbcon
}

func GetMessage(id string) {
	db := connect()

	rows, err := db.Query(`SELECT "Name", "Roll" FROM "Students"`)
	checkError(err)

	defer rows.Close()
	defer db.Close()

	for rows.Next() {
	}
}

func SaveMessage(msg model.Message) sql.Result {
	db := connect()
	if db.Ping() != nil {
		panic(db.Ping().Error())
	}
	insertDynStmt := `insert into Message ("id", "msg") values($1, $2)`
	result, err := db.Exec(insertDynStmt, msg.ID, msg.Msg)
	checkError(err)
	defer db.Close()
	return result
}

func UpdateMessage(msg model.Message) sql.Result {
	db := connect()
	if db.Ping() != nil {
		panic(db.Ping().Error())
	}

	updateDynStmt := `update Message set "msg"=$1 where "id"=$2`
	result, err := db.Exec(updateDynStmt, msg.Msg, msg.ID)

	checkError(err)
	defer db.Close()

	return result
}

func DeleteMessage() {

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
