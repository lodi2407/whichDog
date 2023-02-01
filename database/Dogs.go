package database

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

func Dogs(w http.ResponseWriter, r *http.Request) (rows *sqlx.Rows, err error) {
	rows, err = DatabaseConnexion().Queryx("SELECT * FROM dog ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	return rows, err

}