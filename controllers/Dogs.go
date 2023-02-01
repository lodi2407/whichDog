package controllers

import (
	"net/http"
	"whichDog/database"

	"github.com/jmoiron/sqlx"
)

func Dogs(w http.ResponseWriter, r *http.Request) (rows *sqlx.Rows, err error) {
	rows, err = database.Dogs(w, r)

	return rows, err
}