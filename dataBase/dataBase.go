package dataBase

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DbConnect() *sql.DB {
	connectionString := "user=postgres dbname=ale_loja password=postpass host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	return db
}
