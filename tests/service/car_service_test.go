package service

import (
	"github.com/google/uuid"
	"gopher_rentals/db"
	"gopher_rentals/services"
	"testing"
)

func TestCreateCar(t *testing.T) {
	_ = db.ConfigureDB()

	car1 := map[string]interface{} {
		"model": "Toyota",
		"year": 2011,
		"license_plate": "ABC123ER",
		"current_km": 560.90,
		"max_km": 34.56,
		"fuel_type": "Petrol",
		"hire_price": 908.67,
	}

	car2 := map[string]interface{} {
		"model": "Honda",
		"year": 2021,
		"license_plate": "QWE123TY",
		"current_km": 860.90,
		"max_km": 340.56,
		"fuel_type": "Diesel",
		"hire_price": 1908.67,
	}

	_, err := services.CreateCar(car1)
	if err != nil {
		t.Fatalf("TestCreateCar: error occurred creating car -> %v", err)
	}

	_, err = services.CreateCar(car2)
}

func TestListCars(t *testing.T) {
	_ = db.ConfigureDB()

	cars, err := services.ListCars()
	if err != nil {
		t.Fatalf("TestListCars: error occurred listing cars -> %v", err)
	}

	if len(cars) != 2 {
		t.Fatalf("TestListCars: failed to list correct number of cars")
	}
}

func TestGetCar(t *testing.T) {
	_ = db.ConfigureDB()

	cars, err := services.ListCars()
	if err != nil {
		t.Fatalf("TestGetCar: error occurred listing cars -> %v", err)
	}

	car := cars[0]

	c, err := services.GetCar(car.ID)
	if err != nil {
		t.Fatalf("TestGetCar: error occurred getting car with ID %s", car.ID)
	}

	if c.ID != car.ID {
		t.Fatalf("TestGetCar: car ID %s does not match car ID %s", c.ID, car.ID)
	}
}

func TestUpdateCar(t *testing.T) {
	_ = db.ConfigureDB()

	cars, err := services.ListCars()
	if err != nil {
		t.Fatalf("TestUpdateCar: error occurred listing cars -> %v", err)
	}

	car := cars[0]

	updatedCar := map[string]interface{} {
		"id": car.ID,
		"model": "Mazda",
		"year": 2009,
		"license_plate": "AXC123TY",
		"current_km": 1860.90,
		"max_km": 2340.56,
		"fuel_type": "Gas",
		"hire_price": 108.67,
		"hire_availability": false,
	}

	c, err := services.UpdateCar(updatedCar)
	if err != nil {
		t.Fatalf("TestUpdateCar: error occurred updating car -> %v", err)
	}

	if c.ID != updatedCar["id"].(uuid.UUID) {
		t.Fatalf("TestUpdateCar: updated car ID %s does not match car Id %s",
			updatedCar["id"].(uuid.UUID), c.ID)
	}
}

func TestDeleteCar(t *testing.T)  {
	_ = db.ConfigureDB()

	cars, err := services.ListCars()
	if err != nil {
		t.Fatalf("TestDeleteCar: error occurred listing cars -> %v", err)
	}

	for _, car := range cars {
		err := services.DeleteCar(car.ID)
		if err != nil {
			t.Fatalf("TestDeleteCar: failed to delete car -> %v", err)
		}
	}
}
