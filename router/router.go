package router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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
	r.HandleFunc("/dogs", ListDogs).Methods("GET")
	r.HandleFunc("/dogs/categories", Categories).Methods("GET")
	r.HandleFunc("/dogs/trainability", TrainabilityCategories).Methods("GET")
	r.HandleFunc("/dogs/shedding", SheddingCategories).Methods("GET")
	r.HandleFunc("/dogs/grooming", GroomingCategories).Methods("GET")
	r.HandleFunc("/dogs/energy", EnergyCategories).Methods("GET")
	r.HandleFunc("/dogs/demeanor", DemeanorCategories).Methods("GET")
	r.HandleFunc("/dogs/{id}", GetDog).Methods("GET")
	

	if err := http.ListenAndServe(":3000", r); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello homepage")
}

func ListDogs(w http.ResponseWriter, r *http.Request) {
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

// --------------------------- categories -----------------------------------------------

func Categories(w http.ResponseWriter, r *http.Request) {
	rows, err := controllers.Categories()
	if err != nil {
		return
	}

	var categories []string
	for rows.Next() {
        var categoryName sql.NullString
		if err := rows.Scan(&categoryName); err != nil {
			fmt.Println(err)
			panic(err.Error())
		}

		if categoryName.Valid {
			categories = append(categories, categoryName.String)
		}
    }

	result, err := json.Marshal(categories)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close() 
	defer database.DatabaseConnexion().Close()

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(result); err != nil {
        log.Println("Error writing response:", err)
    }
}

func TrainabilityCategories(w http.ResponseWriter, r *http.Request) {
	trainabilityCategories, err := controllers.TrainabilityCategories()
	if err != nil {
		return
	}

	result, err := json.Marshal(trainabilityCategories)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(result); err != nil {
        log.Println("Error writing response:", err)
    }
}

func SheddingCategories(w http.ResponseWriter, r *http.Request) {
	sheddingCategories, err := controllers.SheddingCategories()
	if err != nil {
		return
	}

	result, err := json.Marshal(sheddingCategories)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(result); err != nil {
        log.Println("Error writing response:", err)
    }
}

func GroomingCategories(w http.ResponseWriter, r *http.Request) {
	groomingCategories, err := controllers.GroomingCategories()
	if err != nil {
		return
	}

	result, err := json.Marshal(groomingCategories)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(result); err != nil {
        log.Println("Error writing response:", err)
    }
}

func EnergyCategories(w http.ResponseWriter, r *http.Request) {
	energyCategories, err := controllers.EnergyCategories()
	if err != nil {
		return
	}

	result, err := json.Marshal(energyCategories)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(result); err != nil {
        log.Println("Error writing response:", err)
    }
}

func DemeanorCategories(w http.ResponseWriter, r *http.Request) {
	demeanorCategories, err := controllers.DemeanorCategories()
	if err != nil {
		return
	}

	result, err := json.Marshal(demeanorCategories)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(result); err != nil {
        log.Println("Error writing response:", err)
    }
}
