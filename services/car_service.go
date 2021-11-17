package services

import (
	"fmt"
	"github.com/google/uuid"
	"gopher_rentals/models"
	"gopher_rentals/repositories"
)

func ListCars() ([]models.Car, error) {
	cars, err := repositories.FindAllCars()
	if err != nil {
		return nil, fmt.Errorf("error fetching all cars")
	}

	for index, car := range cars {
		locations, err := repositories.FindAllLocationsByCar(car.ID)
		if err != nil {
			return nil, fmt.Errorf("error fetching car locations")
		}

		car.Locations = locations
		cars[index] = car
	}

	return  cars, nil
}

func GetCar(id uuid.UUID) (models.Car, error) {
	car, err := repositories.FindCarById(id)
	if err != nil {
		return models.Car{}, fmt.Errorf("no car found with ID %s", id)
	}

	locations, err := repositories.FindAllLocationsByCar(car.ID)
	if err != nil {
		return models.Car{}, fmt.Errorf("error fetching car locations")
	}

	car.Locations = locations

	return car, nil
}

func CreateCar(data map[string]interface{}) (models.Car, error) {
	car := models.Car{
		ID: uuid.New(),
		Model: data["model"].(string),
		Year: data["year"].(int),
		LicensePlate: data["license_plate"].(string),
		CurrentKm: data["current_km"].(float64),
		MaxKm: data["max_kg"].(float64),
		FuelType: data["fuel_type"].(string),
		HirePrice: data["hire_price"].(float64),
		HireAvailability: true,
	}

	_, err := repositories.SaveCar(&car)
	if err != nil {
		return models.Car{}, fmt.Errorf("car was not created")
	}

	return car, nil
}

func UpdateCar(data map[string]interface{}) (models.Car, error) {
	car, err := repositories.FindCarById(data["id"].(uuid.UUID))
	if err != nil {
		return models.Car{}, fmt.Errorf("no car matching ID %s", data["id"])
	}

	car.Model = data["model"].(string)
	car.Year = data["year"].(int)
	car.LicensePlate = data["license_plate"].(string)
	car.CurrentKm = data["current_km"].(float64)
	car.MaxKm = data["max_kg"].(float64)
	car.FuelType = data["fuel_type"].(string)
	car.HirePrice = data["hire_price"].(float64)
	car.HireAvailability = data["hire_availability"].(bool)

	_, err = repositories.UpdateCar(&car)
	if err != nil {
		return models.Car{}, fmt.Errorf("error occurred updating car with ID %s", car.ID)
	}

	return car, nil
}

func DeleteCar(id uuid.UUID) error {
	car, err := repositories.FindCarById(id)
	if err != nil {
		return fmt.Errorf("no car matching ID %s", id)
	}

	_, err = repositories.DeleteCar(car.ID)
	if err != nil {
		return fmt.Errorf("error occurred while deleting customer with ID %s", car.ID)
	}
	
	return nil
}
