package main

import (
	"fmt"
	"gopher_rentals/db"
	"gopher_rentals/routes"
	"log"
)

func main() {
	if err := db.ConfigureDB(); err != nil {
		log.Fatal("Connection could not be made to database", err)
		return
	}

	err := routes.Routes()
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}
}
