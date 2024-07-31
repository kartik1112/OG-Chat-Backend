package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	host := os.Getenv("HOST")
	port, _ := strconv.ParseInt(os.Getenv("PORT"), 10, 32)
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	pgConnString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB, err = sql.Open("postgres", pgConnString)
	if err != nil {
		panic("Could not connect to DB")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
    userId SERIAL,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE PRIMARY KEY,
    passwordHash VARCHAR(255) NOT NULL,
    avatarUrl VARCHAR(255),
    status VARCHAR(255),
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		fmt.Println(err)
		panic("Could not create Table")
	}
}
