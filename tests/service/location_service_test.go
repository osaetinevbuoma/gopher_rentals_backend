package service

import (
	"gopher_rentals/db"
	"gopher_rentals/services"
	"testing"
	"time"
)

func TestCarLocationManagement(t *testing.T) {
	_ = db.ConfigureDB()

	// save car location
	car := map[string]interface{} {
		"model": "Toyota",
		"year": 2011,
		"license_plate": "ABC123ER",
		"current_km": 560.90,
		"max_kg": 34.56,
		"fuel_type": "Petrol",
		"hire_price": 908.67,
	}

	c, err := services.CreateCar(car)
	if err != nil {
		t.Fatalf("TestCarLocationManagement (SaveCarLocation): failed to save car -> %v",
			err)
	}

	location1 := map[string]interface{} {
		"latitude": 78.930,
		"longitude": 289.214,
		"current_location_datetime": time.Now(),
	}

	location2 := map[string]interface{} {
		"latitude": 178.930,
		"longitude": 89.214,
		"current_location_datetime": time.Now(),
	}

	_, err = services.SaveCarLocation(c.ID, location1)
	if err != nil {
		t.Fatalf("TestCarLocationManagement (SaveCarLocation): failed to save car" +
			" location -> %v", err)
	}

	_, err = services.SaveCarLocation(c.ID, location2)

	// get car locations
	locations, err := services.GetCarsRecentLocations(c.ID, 2)
	if err != nil {
		t.Fatalf("TestCarLocationManagement (GetCarLocation): error getting car " +
			"locations -> %v", err)
	}

	if len(locations) != 2 {
		t.Fatalf("TestCarLocationManagement (GetCarLocation): car locations (%d) do not " +
			"match expected number -> %d", len(locations), 2)
	}

	// update car location
	updatedLocation := map[string]interface{} {
		"id": locations[0].ID,
		"latitude": 178.930,
		"longitude": 89.214,
		"current_location_datetime": time.Now(),
	}

	_, err = services.UpdateCarLocation(c.ID, updatedLocation)
	if err != nil {
		t.Fatalf("TestCarLocationManagement (UpdateCarLocation): failed updating car " +
			"location -> %v", err)
	}

	// delete car locations
	for _, location := range locations {
		err = services.DeleteCarLocation(location.ID)
		if err != nil {
			t.Fatalf("TestCarLocationManagement (DeleteCarLocation): failed to delete car" +
				" location -> %v", err)
		}
	}

	locations, err = services.GetCarsRecentLocations(c.ID, 2)
	if err != nil {
		t.Fatalf("TestCarLocationManagement (GetCarLocation): error getting car " +
			"locations 2nd get -> %v", err)
	}

	if len(locations) != 0 {
		t.Fatalf("TestCarLocationManagement (GetCarLocation): car locations (%d) do not " +
			"match expected number second get -> %d", len(locations), 2)
	}
}
