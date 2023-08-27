package database

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DatabaseConnexion() *sqlx.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	  }


	// Open a connection to the database
	db, err := sqlx.Connect("postgres", "postgres://" + os.Getenv("DB_USER")  + ":" + os.Getenv("DB_PASS")  + "@" + os.Getenv("DB_HOST") + "/" + os.Getenv("DB_NAME")  + "?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	return db
}
