package controllers

import (
	"fmt"
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


// --------------------------- categories -----------------------------------------------

func Categories() (rows *sqlx.Rows, err error) {
	rows, err = database.Categories()

	return rows, err
}

func TrainabilityCategories() (trainabilityCategories []string, err error) {
	rows, err := database.TrainabilityCategories()

	for rows.Next() {
        var categoryName sql.NullString
		if err := rows.Scan(&categoryName); err != nil {
			fmt.Println(err)
			panic(err.Error())
		}

		if categoryName.Valid {
			trainabilityCategories = append(trainabilityCategories, categoryName.String)
		}
    }

	defer rows.Close() 
	defer database.DatabaseConnexion().Close()

	return trainabilityCategories, err
}

func SheddingCategories() (sheddingCategories []string, err error) {
	rows, err := database.SheddingCategories()

	for rows.Next() {
        var categoryName sql.NullString
		if err := rows.Scan(&categoryName); err != nil {
			fmt.Println(err)
			panic(err.Error())
		}

		if categoryName.Valid {
			sheddingCategories = append(sheddingCategories, categoryName.String)
		}
    }
	
	defer rows.Close() 
	defer database.DatabaseConnexion().Close()

	return sheddingCategories, err
}

func GroomingCategories() (groomingCategories []string, err error) {
	rows, err := database.GroomingCategories()

	for rows.Next() {
        var categoryName sql.NullString
		if err := rows.Scan(&categoryName); err != nil {
			fmt.Println(err)
			panic(err.Error())
		}

		if categoryName.Valid {
			groomingCategories = append(groomingCategories, categoryName.String)
		}
    }
	
	defer rows.Close() 
	defer database.DatabaseConnexion().Close()

	return groomingCategories, err
}

func EnergyCategories() (energyCategories []string, err error) {
	rows, err := database.EnergyCategories()

	for rows.Next() {
        var categoryName sql.NullString
		if err := rows.Scan(&categoryName); err != nil {
			fmt.Println(err)
			panic(err.Error())
		}

		if categoryName.Valid {
			energyCategories = append(energyCategories, categoryName.String)
		}
    }
	
	defer rows.Close() 
	defer database.DatabaseConnexion().Close()

	return energyCategories, err
}

func DemeanorCategories() (demeanorCategories []string, err error) {
	rows, err := database.DemeanorCategories()

	for rows.Next() {
        var categoryName sql.NullString
		if err := rows.Scan(&categoryName); err != nil {
			fmt.Println(err)
			panic(err.Error())
		}

		if categoryName.Valid {
			demeanorCategories = append(demeanorCategories, categoryName.String)
		}
    }
	
	defer rows.Close() 
	defer database.DatabaseConnexion().Close()

	return demeanorCategories, err
}
