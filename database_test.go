package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestDatabase(t *testing.T) {
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
	defer db.Close()
}

func TestSelect(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT * FROM member"
	rows, err := db.QueryContext(ctx,query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int32
		var name, job, hobby string
		err := rows.Scan(&id, &name, &job, &hobby)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID :",id)
		fmt.Println("Name :",name)
		fmt.Println("Job :",job)
		fmt.Println("hobby :",hobby)
		defer rows.Close()
	}
}


