package repositories

import (
	"fmt"
	"gopher_rentals/db"
	"gopher_rentals/models"
)

func SaveCustomerHireCar(data *models.CustomerHireCar) (int64, error) {
	result, err := db.DB.Exec("INSERT INTO customer_hire_car (id, customer_id, car_id, " +
		"hire_date, return_date) VALUES (?, ?, ?, ?, ?) ", data.ID, data.Customer.ID, data.Car.ID,
		data.HireDate, data.ReturnDate)
	if err != nil {
		return 0, fmt.Errorf("SaveCustomerHireCar: %v", err)
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("SaveCustomerHireCar: %v", err)
	}

	return row, nil
}

func UpdateCustomerHireCar(data *models.CustomerHireCar) (int64, error) {
	result, err := db.DB.Exec("UPDATE customer_hire_car SET customer_id = ?, car_id = ?, " +
		"hire_date = ?, return_date = ? WHERE id = ?", data.Customer.ID, data.Car.ID, data.HireDate,
		data.ReturnDate, data.ID)
	if err != nil {
		return 0, fmt.Errorf("UpdateCustomerHireCar: %v", err)
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("UpdateCustomerHireCar: %v", err)
	}

	return row, nil
}
