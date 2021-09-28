package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

func CreateConnection() *sql.DB {
	//load env file
	err := godotenv.Load(".env")

	if err!=nil {
		log.Fatalf("Error loading .env file")
	}

	//Connect to postgres db
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err!=nil {
		panic(err)
	}

	//Create table if not exist
	sqlQuery := "CREATE TABLE IF NOT EXISTS movies (id serial primary key, title varchar, rating integer);"
	db.Query(sqlQuery)

	//Check Connection
	err = db.Ping()

	if err!=nil {
		panic(err)
	}

	fmt.Println("Connected to db successfully")

	return db
}