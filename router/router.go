package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"whichDog/controllers"
	"whichDog/database"
	"whichDog/models"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Router() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomePage)
	r.HandleFunc("/dogs", Dogs).Methods("GET")
	r.HandleFunc("/dogs/{id}", GetDog).Methods("GET")

	http.ListenAndServe(":3000", r)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello homepage")
}

func Dogs(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	rows, err := controllers.ListDogs(queryParams)
	if err != nil {
		fmt.Println("err", err)
	}

	dogs := []models.Dog{}
	err = sqlx.StructScan(rows, &dogs)
	if err != nil {
		fmt.Println("err", err)
	}
		
	results, err := json.Marshal(dogs)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	defer rows.Close()
	defer database.DatabaseConnexion().Close()

	w.Header().Set("Content-Type", "application/json")
	w.Write(results)
}

func GetDog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id := vars["id"]
	rows, err := controllers.GetDog(id)
	if err != nil {
		return
	}

	dog := models.Dog{}
	for rows.Next() {
	err = rows.StructScan(&dog)
		if err != nil {
			fmt.Println("err", err)
		}
	}
		
	result, err := json.Marshal(dog)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	defer rows.Close()
	defer database.DatabaseConnexion().Close()

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}