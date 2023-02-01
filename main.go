package main

import (
	"whichDog/database"
	"whichDog/router"
)

func main() {
	router.Router()
	database.DatabaseConnexion()
}