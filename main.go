package main

import (
	"gopher_rentals/db"
	"log"
)

func main() {
	if err := db.ConfigureDB(); err != nil {
		log.Fatal("Connection could not be made to database", err)
		return
	}
}
