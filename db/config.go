package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	StorageDB *sql.DB
)

func NewConnection() (*sql.DB, error) {
	// Get the enviroment vars
	envErr := godotenv.Load("../.env")
	if envErr != nil {
		log.Fatal("error in load .env file")
	}
	dbconnection := os.Getenv("DB_CONNECTION")
	if dbconnection == "" {
		panic("error reading enviroment value: db_connection")
	}
	var err error
	StorageDB, err = sql.Open("mysql", dbconnection)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	return StorageDB, nil
}
