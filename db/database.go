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

func GetMessage(id string) []model.Message {
	db := connect()
	selectDynStmt := `SELECT id, msg FROM Message where "id"=$1`
	rows, err := db.Query(selectDynStmt, id)

	messages := []model.Message{}

	checkError(err)
	defer db.Close()
	defer rows.Close()

	for rows.Next() {
		var msg model.Message
		if err := rows.Scan(&msg.ID, &msg.Msg); err != nil {
			panic(err)
		}

		messages = append(messages, msg)
	}

	return messages
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

func DeleteMessage(id string) sql.Result {
	db := connect()
	if db.Ping() != nil {
		panic(db.Ping().Error())
	}

	deleteDynStmt := `delete from Message where "id"=$1`
	result, err := db.Exec(deleteDynStmt, id)
	checkError(err)

	defer db.Close()

	return result
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}

}
