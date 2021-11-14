package repositories

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"gopher_rentals/db"
	"gopher_rentals/models"
)

func FindAllCars() ([]models.Car, error) {
	var cars []models.Car

	rows, err := db.DB.Query("SELECT * FROM cars")
	if err != nil {
		return nil, fmt.Errorf("FindAllCars: %v", err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.ID, &car.Model, &car.Year, &car.LicensePlate, &car.CurrentKm,
			&car.MaxKm, &car.FuelType, &car.HirePrice, &car.HireAvailability); err != nil {
			return nil, fmt.Errorf("FindAllCars: %v", err)
		}
		cars = append(cars, car)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FindAllCars: %v", err)
	}

	return cars, nil
}

func FindCarById(id uuid.UUID) (models.Car, error) {
	var car models.Car

	row := db.DB.QueryRow("SELECT * FROM cars WHERE id = ?", id)
	if err := row.Scan(&car.ID, &car.Model, &car.Year, &car.LicensePlate, &car.CurrentKm,
		&car.MaxKm, &car.FuelType, &car.HirePrice, &car.HireAvailability); err != nil {
		if err == sql.ErrNoRows {
			return car, fmt.Errorf("FindCarById: No user with id %s", id)
		}

		return car, fmt.Errorf("FindCarById: %v", err)
	}

	return car, nil
}

func SaveCar(car models.Car) (int64, error) {
	result, err := db.DB.Exec("INSERT INTO cars " +
		"(id, model, year, license_plate, current_km, max_km, fuel_type, hire_price, hire_availability) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", car.ID, car.Model, car.Year, car.LicensePlate,
		car.CurrentKm, car.MaxKm, car.FuelType, car.HirePrice, car.HireAvailability)
	if err != nil {
		return 0, fmt.Errorf("SaveCar: %v", err)
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("SaveCar: %v", err)
	}

	return row, nil
}

func UpdateCar(car models.Car) (int64, error) {
	result, err := db.DB.Exec("UPDATE cars SET model = ?, year = ?,  license_plate = ?" +
		"current_km = ?, max_km = ?, fuel_type = ?, hire_price = ?, hire_availability = ? " +
		"WHERE id = ?", car.Model, car.Year, car.LicensePlate, car.CurrentKm, car.MaxKm,
		car.FuelType, car.HirePrice, car.HireAvailability, car.ID)
	if err != nil {
		return 0, fmt.Errorf("UpdateCar: %v", err)
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("UpdateCar: %v", err)
	}

	return row, nil
}

func DeleteCar(id string) (int64, error) {
	result, err := db.DB.Exec("DELETE FROM cars WHERE id = ?", id)
	if err != nil {
		return 0, fmt.Errorf("DeleteCar: %v", err)
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("DeleteCar: %v", err)
	}

	return row, nil
}
