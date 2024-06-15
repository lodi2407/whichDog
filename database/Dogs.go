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
		fmt.Println(full)
		fmt.Println(args)
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

func ColumnsName() (columnNames []string) {
	query := "SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = N'dog'"
	rows, err := DatabaseConnexion().Queryx(query)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	 for rows.Next() {
        var columnName string
        if err := rows.Scan(&columnName); err != nil {
			fmt.Println(err)
			panic(err.Error())
        }
        columnNames = append(columnNames, columnName)
    }

	defer rows.Close()

	return columnNames
}