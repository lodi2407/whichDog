package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func DatabaseConnexion() *sqlx.DB {
	// Open a connection to the database
	db, err := sqlx.Connect("mysql", "root:@tcp(localhost:3306)/dogs")
	if err != nil {
		panic(err.Error())
	}

	return db
}
