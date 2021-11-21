package service

import (
	"gopher_rentals/db"
	"gopher_rentals/models"
	"gopher_rentals/services"
	"testing"
)

func TestCreateCar(t *testing.T) {
	_ = db.ConfigureDB()

	car1 := models.Car{
		Model: "Toyota",
		Year: 2011,
		LicensePlate: "ABC123ER",
		CurrentKm: 560.90,
		MaxKg: 34.56,
		FuelType: "Petrol",
		HirePrice: 908.67,
	}

	car2 := models.Car {
		Model: "Honda",
		Year: 2021,
		LicensePlate: "QWE123TY",
		CurrentKm: 860.90,
		MaxKg: 340.56,
		FuelType: "Diesel",
		HirePrice: 1908.67,
	}

	_, err := services.CreateCar(&car1)
	if err != nil {
		t.Fatalf("TestCreateCar: error occurred creating car -> %v", err)
	}

	_, err = services.CreateCar(&car2)
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
	car.Model = "Mazda"
	car.Year = 2009
	car.LicensePlate = "AXC123TY"
	car.CurrentKm =  1860.90
	car.MaxKg = 2340.56
	car.FuelType = "Gas"
	car.HirePrice = 108.67
	car.HireAvailability = false

	c, err := services.UpdateCar(&car)
	if err != nil {
		t.Fatalf("TestUpdateCar: error occurred updating car -> %v", err)
	}

	if c.ID != car.ID {
		t.Fatalf("TestUpdateCar: updated car ID %s does not match car Id %s", car.ID, c.ID)
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
