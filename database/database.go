package database

import (
	"database/sql"
	"fmt"
	"log"
)

func OpenConnection(host string, port int, user string, password string, dbname string) *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	))
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}
