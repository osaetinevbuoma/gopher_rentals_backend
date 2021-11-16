package repositories

import (
	"github.com/google/uuid"
	"gopher_rentals/db"
	"gopher_rentals/models"
	"gopher_rentals/repositories"
	"testing"
)

func TestSaveCar(t *testing.T) {
	if err := db.ConfigureDB(); err != nil {
		t.Fatalf("Connection established successfully")
	}

	car := models.Car{
		ID:               uuid.New(),
		Model:            "Toyota",
		Year:             2009,
		LicensePlate:     "ABC123ER",
		CurrentKm:        1000,
		MaxKg:            50,
		FuelType:         "Petrol",
		HirePrice:        5000,
		HireAvailability: true,
	}

	row, err := repositories.SaveCar(&car)
	if err != nil {
		t.Fatalf("TestSaveCar not saved: %v", err)
	}

	if row != 1 {
		t.Fatalf("TestSaveCar row not inserted")
	}
}

func TestFindAllCars(t *testing.T) {
	if err := db.ConfigureDB(); err != nil {
		t.Fatalf("Connection established successfully")
	}

	cars, err := repositories.FindAllCars()
	if err != nil {
		t.Fatalf("TestFindAllCars error: %v", err)
	}

	if len(cars) == 0 {
		t.Fatalf("TestFindAllCars length is 0")
	}
}

func TestFindCarUpdateCarDeleteCar(t *testing.T) {
	if err := db.ConfigureDB(); err != nil {
		t.Fatalf("Connection established successfully")
	}

	cars, err := repositories.FindAllCars()
	if err != nil {
		t.Fatalf("TestFindCarUpdateCarDeleteCar error: %v", err)
	}

	if len(cars) == 0 {
		t.Fatalf("TestFindCarUpdateCarDeleteCar length is 0")
	}

	carId := cars[0].ID
	car, err := repositories.FindCarById(carId)
	if err != nil {
		t.Fatalf("TestFindCarUpdateCarDeleteCar error: %v", err)
	}

	if car.ID != carId {
		t.Fatalf("TestFindCarUpdateCarDeleteCar: Car IDs do not match")
	}

	car.FuelType = "Diesel"
	car.Year = 2011

	row, err := repositories.UpdateCar(&car)
	if err != nil {
		t.Fatalf("TestFindCarUpdateCarDeleteCar error: %v", err)
	}

	if row != 1 {
		t.Fatalf("TestFindCarUpdateCarDeleteCar: update row is not 1")
	}

	row, err = repositories.DeleteCar(car.ID)
	if err != nil {
		t.Fatalf("TestFindCarUpdateCarDeleteCar error: %v", err)
	}

	if row != 1 {
		t.Fatalf("TestFindCarUpdateCarDeleteCar: delete row is not 1")
	}
}
