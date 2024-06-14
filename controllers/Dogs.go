package controllers

import (
	"net/url"
	"whichDog/database"

	"github.com/jmoiron/sqlx"
)

func Dogs( queryParams url.Values) (rows *sqlx.Rows, err error) {
	category := queryParams.Get("category")
	// sheddingCategory := queryParams.Get("sheddingCategory")
	
	var query string
	args := []interface{}{}

	if category != "" {
		query += " AND category = $1"
		args = append(args, category)
	}

	// if sheddingCategory != "" {
	// 	query += " AND shedding_category = $2"
	// 	args = append(args, category)
	// }

	rows, err = database.Dogs(query, args)

	return rows, err
}

func GetDog(id string) (rows *sqlx.Rows, err error) {
	rows, err = database.GetDog(id)

	return rows, err
}