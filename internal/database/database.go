package database

import (
	"database/sql"
	"fmt"

	_ "github.com/sijms/go-ora/v2"
)

var dbConfig = map[string]string{
	"password": "password",
	"port": "1521",
	"server": "localhost",
	"service": "XE",
	"username": "SYSTEM",
}

var (
	DBCon *sql.DB
)


func Connect()(*sql.DB, error) {
	connectionString := "oracle://" + dbConfig["username"] + ":" + dbConfig["password"] + "@" + dbConfig["server"] + ":" + dbConfig["port"] + "/" + dbConfig["service"]
	db, err := sql.Open("oracle", connectionString)

	if err != nil {
		fmt.Println("Error while connecting to database", err)
	}

	_ = db.Ping()

	return db,err
}