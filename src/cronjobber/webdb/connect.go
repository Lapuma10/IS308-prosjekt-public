package webdb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "BSBACIT2020"
	hostname = "web-db"
	dbname   = "project"
)

func constBuilder() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func Connect() *sql.DB {
	db, err := sql.Open("mysql", constBuilder())

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return db
}
