package database

import (
	"database/sql"
	"log"
)

var DbConn *sql.DB

func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/inventory")
	if err != nil {
		log.Fatal(err)
	}

}
