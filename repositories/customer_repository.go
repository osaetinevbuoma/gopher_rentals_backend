package repositories

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"gopher_rentals/db"
	"gopher_rentals/models"
)

func FindAllCustomers() ([]models.Customer, error)  {
	var customers []models.Customer
	
	rows, err := db.DB.Query("SELECT * FROM customers")
	if err != nil {
		return nil, fmt.Errorf("FindAllCustomers: %v", err)
	}
	
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var customer models.Customer
		if err := rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName,
			&customer.Nationality, &customer.IdentificationNumber,
			&customer.IdentificationType); err != nil {
			return nil, fmt.Errorf("FindAllCustomers: %v", err)
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FindAllCustomers: %v", err)
	}

	return customers, nil
}

func FindCustomerById(id uuid.UUID) (models.Customer, error) {
	var customer models.Customer

	row := db.DB.QueryRow("SELECT * FROM customers WHERE id = ?", id)
	if err := row.Scan(&customer.ID, &customer.FirstName, &customer.LastName,
		&customer.Nationality, &customer.IdentificationNumber,
		&customer.IdentificationType); err != nil {
		if err == sql.ErrNoRows {
			return customer, fmt.Errorf("FindCustomerById: No user with id %s", id)
		}

		return customer, fmt.Errorf("FindCustomerById: %v", err)
	}

	return customer, nil
}

func SaveCustomer(customer *models.Customer) (int64, error) {
	result, err := db.DB.Exec("INSERT INTO customers " +
		"(id ,first_name, last_name, nationality, identification_number, identification_type) VALUES " +
		"(?, ?, ?, ?, ?, ?)", customer.ID, customer.FirstName, customer.LastName,
		customer.Nationality, customer.IdentificationNumber, customer.IdentificationType)
	if err != nil {
		return 0, fmt.Errorf("SaveCustomer: %v", err)
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("SaveCustomer: %v", err)
	}

	return row, nil
}

func UpdateCustomer(customer *models.Customer) (int64, error) {
	result, err := db.DB.Exec("UPDATE customers SET first_name = ?, last_name = ?, " +
		"nationality = ?, identification_number = ?, identification_type = ? WHERE id = ?",
		customer.FirstName, customer.LastName, customer.Nationality, customer.IdentificationNumber,
		customer.IdentificationType, customer.ID)
	if err != nil {
		return 0, fmt.Errorf("UpdateCustomer: %v", err)
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("UpdateCustomer: %v", err)
	}

	return row, nil
}

func DeleteCustomer(id uuid.UUID) (int64, error) {
	result, err := db.DB.Exec("DELETE FROM customers WHERE id = ?", id)
	if err != nil {
		return 0, fmt.Errorf("DeleteCustomer: %v", err)
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("DeleteCustomer: %v", err)
	}

	return row, nil
}
