package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Dogs(query string, args []interface{}) (rows *sqlx.Rows, err error) {
	if query == "" {
		rows, err = DatabaseConnexion().Queryx("SELECT * FROM dog ORDER BY id ASC")
		if err != nil {
			panic(err.Error())
		}
	} else {
		full := "SELECT * FROM dog WHERE 1=1"
		full += query
		rows, err = DatabaseConnexion().Queryx(full, args...)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
	}

	return rows, err

}

func GetDog(id string) (rows *sqlx.Rows, err error) {
	rows, err = DatabaseConnexion().Queryx("SELECT * FROM dog WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	return rows, err

}