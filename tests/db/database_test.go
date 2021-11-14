package db

import (
	"gopher_rentals/db"
	"testing"
)

func TestDBConnection(t *testing.T) {
	if err := db.ConfigureDB(); err != nil {
		t.Fatalf("Connection established successfully")
	}
}
