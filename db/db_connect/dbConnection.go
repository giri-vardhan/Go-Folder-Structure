package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "giriopen"
	dbname   = "giri"
)


var DB = ConnectDB()

func ConnectDB() *sql.DB {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	
	DB, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	err = DB.Ping()
	CheckError(err)
	fmt.Println("successfully connected")
	return DB
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
