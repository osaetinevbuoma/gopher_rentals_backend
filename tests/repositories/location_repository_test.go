package repositories

import (
	"github.com/google/uuid"
	"gopher_rentals/db"
	"gopher_rentals/models"
	"gopher_rentals/repositories"
	"testing"
	"time"
)

func TestLocationManagement(t *testing.T) {
	_ = db.ConfigureDB()

	// We need a car to assign locations
	// To be deleted at the end of test
	car := models.Car{
		ID: uuid.New(),
		Model: "Toyota",
		Year: 2009,
		LicensePlate: "ABC123ER",
		CurrentKm: 1000,
		MaxKm: 50,
		FuelType: "Petrol",
		HirePrice: 5000,
		HireAvailability: true,
	}

	row, _ := repositories.SaveCar(car)

	if row != 1 {
		t.Fatalf("TestLocationManagement: SaveCar -> row not inserted")
	}

	location := models.Location{
		ID: uuid.New(),
		Car: car,
		Latitude: 12.123,
		Longitude: 5.902,
		CurrentLocationDatetime: time.Now(),
	}

	row, _ = repositories.SaveLocation(location)
	if row != 1 {
		t.Fatalf("TestLocationManagement: SaveLocation -> row not inserted")
	}

	// testing finding locations by car
	locations, err := repositories.FindAllLocationsByCar(car.ID)
	if err != nil {
		t.Fatalf("TestLocationManagement: FindAllLocationsByCar -> %v", err)
	}

	if len(locations) == 0 {
		t.Fatalf("TestLocationManagement: FindAllLocationsByCar -> equals 0")
	}

	location2 := models.Location{
		ID: uuid.New(),
		Car: car,
		Latitude: 14.920,
		Longitude: 9.094,
		CurrentLocationDatetime: time.Now(),
	}

	// testing finding filtered locations
	row, _ = repositories.SaveLocation(location2)
	if row != 1 {
		t.Fatalf("TestLocationManagement: SaveLocation -> row not inserted")
	}

	locations, err = repositories.FindLocationsByCarFiltered(car.ID, 2)
	if err != nil {
		t.Fatalf("TestLocationManagement: FindLocationsByCarFiltered -> %v", err)
	}

	if len(locations) != 2 {
		t.Fatalf("TestLocationManagement: FindLocationsByCarFiltered -> not equals 2")
	}

	// testing updating a location
	location.CurrentLocationDatetime = time.Now()
	location.Longitude = 10.903
	location.Longitude = 8.424

	row, err = repositories.UpdateLocation(location)
	if err != nil {
		t.Fatalf("TestLocationManagement: UpdateLocation -> %v", err)
	}

	if row != 1 {
		t.Fatalf("TestLocationManagement: UpdateLocation -> row not updated")
	}

	// testing deleting location
	row, err = repositories.DeleteLocation(location.ID)
	if err != nil {
		t.Fatalf("TestLocationManagement: DeleteLocation -> %v", err)
	}

	if row != 1 {
		t.Fatalf("TestLocationManagement: DeleteLocation -> row not deleted")
	}

	// clean DB
	_,_ = repositories.DeleteLocation(location2.ID)
	_,_ = repositories.DeleteCar(car.ID)
}
