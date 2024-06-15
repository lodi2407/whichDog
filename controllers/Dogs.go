package controllers

import (
	"net/url"
	"strconv"
	"whichDog/database"

	"github.com/jmoiron/sqlx"
)

func ListDogs(queryParams url.Values) (rows *sqlx.Rows, err error) {
	columnNames := database.ColumnsName()
	
	var query string
	args := []interface{}{}
	paramIndex := 1

	for _, column := range columnNames {
		if value := queryParams.Get(column); value != "" {
			query += " AND " + column + " = $" + strconv.Itoa(paramIndex)
			args = append(args, value)
			paramIndex++
		}
	}

	rows, err = database.Dogs(query, args)

	return rows, err
}

func GetDog(id string) (rows *sqlx.Rows, err error) {
	rows, err = database.GetDog(id)

	return rows, err
}