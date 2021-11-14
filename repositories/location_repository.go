package repositories

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"gopher_rentals/db"
	"gopher_rentals/models"
)

func FindAllLocationsByCar(carId uuid.UUID) ([]models.Location, error) {
	var locations []models.Location

	rows, err := db.DB.Query("SELECT * FROM locations WHERE car_id = ?", carId)
	if err != nil {
		return nil, fmt.Errorf("FindAllLocationsByCar: %v", err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var location models.Location

		car, err := FindCarById(carId)
		if err != nil {
			return nil, fmt.Errorf("FindAllLocationsByCar: %v", err)
		}

		if err := rows.Scan(&location.ID, &car.ID, &location.Latitude, &location.Longitude,
			&location.CurrentLocationDatetime); err != nil {
			return nil, fmt.Errorf("FindAllLocationsByCar: %v", err)
		}
		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FindAllLocationsByCar: %v", err)
	}

	return locations, nil
}

func FindLocationsByCarFiltered(carId uuid.UUID, limit int) ([]models.Location, error) {
	var locations []models.Location

	rows, err := db.DB.Query("SELECT * FROM locations WHERE car_id = ? LIMIT ?", carId, limit)
	if err != nil {
		return nil, fmt.Errorf("FindLocationsByCarFiltered: %v", err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var location models.Location

		car, err := FindCarById(carId)
		if err != nil {
			return nil, fmt.Errorf("FindLocationsByCarFiltered: %v", err)
		}

		if err := rows.Scan(&location.ID, &car.ID, &location.Latitude, &location.Longitude,
			&location.CurrentLocationDatetime); err != nil {
			return nil, fmt.Errorf("FindLocationsByCarFiltered: %v", err)
		}
		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FindLocationsByCarFiltered: %v", err)
	}

	return locations, nil
}

func SaveLocation(location models.Location) (int64, error) {
	result, err := db.DB.Exec("INSERT INTO locations (id, car_id, latitude, longitude, " +
		"current_location_datetime) VALUES (?, ?, ?, ?, ?)", location.ID, location.Car.ID,
		location.Latitude, location.Longitude, location.CurrentLocationDatetime)
	if err != nil {
		return 0, fmt.Errorf("SaveLocation: %v", err)
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("SaveLocation: %v", err)
	}

	return row, nil
}

func UpdateLocation(location models.Location) (int64, error) {
	result, err := db.DB.Exec("UPDATE locations SET car_id = ?, latitude = ?,  longitude = ?, " +
		"current_location_datetime = ? WHERE id = ?", location.Car.ID, location.Latitude,
		location.Longitude, location.CurrentLocationDatetime, location.ID)
	if err != nil {
		return 0, fmt.Errorf("UpdateLocation: %v", err)
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("UpdateLocation: %v", err)
	}

	return row, nil
}

func DeleteLocation(id uuid.UUID) (int64, error) {
	result, err := db.DB.Exec("DELETE FROM locations WHERE id = ?", id)
	if err != nil {
		return 0, fmt.Errorf("DeleteLocation: %v", err)
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("DeleteLocation: %v", err)
	}

	return row, nil
}
