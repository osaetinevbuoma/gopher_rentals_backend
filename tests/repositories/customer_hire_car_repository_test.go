package repositories

import (
	"github.com/google/uuid"
	"gopher_rentals/db"
	"gopher_rentals/models"
	"gopher_rentals/repositories"
	"testing"
	"time"
)

func TestSaveCustomerHireCar(t *testing.T) {
	_ = db.ConfigureDB()

	car := models.Car{
		ID:               uuid.New(),
		Model:            "Toyota",
		Year:             2009,
		LicensePlate:     "ABC123ER",
		CurrentKm:        1000,
		MaxKm:            50,
		FuelType:         "Petrol",
		HirePrice:        5000,
		HireAvailability: true,
	}

	customer := models.Customer{
		ID:                   uuid.New(),
		FirstName:            "John",
		LastName:             "Doe",
		Nationality:          "Nigerian",
		IdentificationNumber: "ABC123",
		IdentificationType:   "International Passport",
	}

	_, _ = repositories.SaveCar(car)
	_, _ = repositories.SaveCustomer(customer)

	assignCarToCustomer := models.CustomerHireCar{
		ID:         uuid.New(),
		Customer:   customer,
		Car:        car,
		HireDate:   time.Now(),
		ReturnDate: time.Now().Add(time.Duration(time.Sunday)),
	}

	row, err := repositories.SaveCustomerHireCar(assignCarToCustomer)
	if err != nil {
		t.Fatalf("TestSaveCustomerHireCar -> %v", err)
	}

	if row == 0 {
		t.Fatalf("TestSaveCustomerHireCar -> assignment was inserted")
	}

	_, _ = repositories.DeleteCar(car.ID)
	_, _ = repositories.DeleteCustomer(customer.ID)
}

func TestUpdateCustomerHireCar(t *testing.T) {
	_ = db.ConfigureDB()

	car := models.Car{
		ID:               uuid.New(),
		Model:            "Toyota",
		Year:             2010,
		LicensePlate:     "ABCD123ER",
		CurrentKm:        2000,
		MaxKm:            100,
		FuelType:         "Diesel",
		HirePrice:        5000,
		HireAvailability: true,
	}

	customer := models.Customer{
		ID:                   uuid.New(),
		FirstName:            "John 3",
		LastName:             "Doe 3",
		Nationality:          "Nigerian",
		IdentificationNumber: "ABCD123",
		IdentificationType:   "International Passport",
	}

	_, _ = repositories.SaveCar(car)
	_, _ = repositories.SaveCustomer(customer)

	assignCarToCustomer := models.CustomerHireCar{
		ID:         uuid.New(),
		Customer:   customer,
		Car:        car,
		HireDate:   time.Now(),
		ReturnDate: time.Now().Add(time.Duration(time.Sunday)),
	}

	_, _ = repositories.SaveCustomerHireCar(assignCarToCustomer)

	// update
	assignCarToCustomer.HireDate = time.Now()
	assignCarToCustomer.ReturnDate = time.Now().Add(time.Hour * 720)

	row, err := repositories.UpdateCustomerHireCar(assignCarToCustomer)
	if err != nil {
		t.Fatalf("TestUpdateCustomerHireCar -> %v", err)
	}

	if row == 0 {
		t.Fatalf("TestUpdateCustomerHireCar -> assignment was not updated")
	}

	_, _ = repositories.DeleteCar(car.ID)
	_, _ = repositories.DeleteCustomer(customer.ID)
}
