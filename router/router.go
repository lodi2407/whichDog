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

	http.ListenAndServe(":3000", r)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello homepage")
}

func Dogs(w http.ResponseWriter, r *http.Request) {
	rows, err := controllers.Dogs(w, r)

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