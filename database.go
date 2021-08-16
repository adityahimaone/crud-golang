package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func GetConnection() *sql.DB {
	dbDriver := "mysql"
	dbUsername := "root"
	// dbPassword := ""
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "projectone"
	db, err := sql.Open(dbDriver, dbUsername+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
