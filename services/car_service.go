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

func CreateCar(car *models.Car) (*models.Car, error) {
	car.ID = uuid.New()
	_, err := repositories.SaveCar(car)
	if err != nil {
		return nil, fmt.Errorf("car was not created")
	}

	return car, nil
}

func UpdateCar(car *models.Car) (*models.Car, error) {
	_, err := repositories.FindCarById(car.ID)
	if err != nil {
		return nil, fmt.Errorf("no car matching ID %s", car.ID)
	}

	_, err = repositories.UpdateCar(car)
	if err != nil {
		return nil, fmt.Errorf("error occurred updating car with ID %s", car.ID)
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
