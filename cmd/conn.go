package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	go_ora "github.com/sijms/go-ora/v2"
)

func conn() *sqlx.DB {
	username := "SYS"
	password := "myP@ssw0rd"
	host := "localhost"
	port := 1521
	serviceName := "LPMDB"

	connStr := go_ora.BuildUrl(host, port, serviceName, username, password, nil)
	db, err := sqlx.Open("oracle", connStr)

	if err != nil {
		log.Fatal("Connection Error: ", err)
	}

	return db
}
