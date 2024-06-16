package database

import (
	"api/src/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	connect, err := sql.Open("postgres", config.DbConnectionString)
	if err != nil || connect == nil {
		log.Fatal("failed to open conection", err)
	}

	if err := connect.Ping(); err != nil {
		connect.Close()
		log.Fatal("failed to ping on database", err)
	}

	return connect
}
