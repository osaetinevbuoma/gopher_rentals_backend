package services

import (
	"fmt"
	"github.com/google/uuid"
	"gopher_rentals/models"
	"gopher_rentals/repositories"
	"time"
)

func GetCarLocations(carId uuid.UUID) ([]models.Location, error) {
	_, err := repositories.FindCarById(carId)
	if err != nil {
		return nil, fmt.Errorf("no car matching ID %s", carId)
	}

	locations, err := repositories.FindAllLocationsByCar(carId)
	if err != nil {
		return nil, fmt.Errorf("locations for %s could not be fetched", carId)
	}

	return locations, nil
}

func GetCarsRecentLocations(carId uuid.UUID, recent int) ([]models.Location, error) {
	car, err := repositories.FindCarById(carId)
	if err != nil {
		return nil, fmt.Errorf("no car matching ID %s", carId)
	}

	locations, err := repositories.FindLocationsByCarFiltered(car.ID, recent)
	if err != nil {
		return nil, fmt.Errorf("error occurred fetching locations for car with ID %s", carId)
	}

	return locations, nil
}

func SaveCarLocation(carId uuid.UUID, data map[string]interface{}) (models.Location, error) {
	car, err := repositories.FindCarById(carId)
	if err != nil {
		return models.Location{}, fmt.Errorf("no car matching ID %s", carId)
	}

	currentLocationDatetime, err := time.Parse("2006-01-02 15:04:05",
		data["current_location_datetime"].(string))
	if err != nil {
		return models.Location{}, err
	}

	location := models.Location{
		ID: uuid.New(),
		Car: car,
		Latitude: data["latitude"].(float64),
		Longitude: data["longitude"].(float64),
		CurrentLocationDatetime: currentLocationDatetime,
	}

	_, err = repositories.SaveLocation(&location)
	if err != nil {
		return models.Location{}, fmt.Errorf("car location was not saved")
	}

	return location, nil
}

func UpdateCarLocation(carId uuid.UUID, data map[string]interface{}) (models.Location, error) {
	car, err := repositories.FindCarById(carId)
	if err != nil {
		return models.Location{}, fmt.Errorf("no car matching ID %s", carId)
	}

	location, err := repositories.FindLocationById(uuid.MustParse(data["id"].(string)))
	if err != nil {
		return models.Location{}, fmt.Errorf("no record matching location with ID %s",
			data["id"].(uuid.UUID))
	}

	currentLocationDatetime, err := time.Parse("2006-01-02 15:04:05",
		data["current_location_datetime"].(string))
	if err != nil {
		return models.Location{}, fmt.Errorf("error converting datetime value")
	}

	location.Car = car
	location.CurrentLocationDatetime = currentLocationDatetime
	location.Latitude = data["latitude"].(float64)
	location.Longitude = data["longitude"].(float64)

	_, err = repositories.UpdateLocation(&location)
	if err != nil {
		return models.Location{}, fmt.Errorf("error occurred updating location")
	}

	return location, nil
}

func DeleteCarLocation(id uuid.UUID) error {
	_, err := repositories.DeleteLocation(id)
	if err != nil {
		return fmt.Errorf("error occurred deleting location %v", err)
	}

	return nil
}
