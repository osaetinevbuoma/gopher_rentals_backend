package services

import (
	"fmt"
	"github.com/google/uuid"
	"gopher_rentals/models"
	"gopher_rentals/repositories"
)

func ListCustomers() ([]models.Customer, error) {
	customers, err := repositories.FindAllCustomers()
	if err != nil {
		return nil, fmt.Errorf("error fetching all customers")
	}

	return customers, nil
}

func GetCustomer(id uuid.UUID) (models.Customer, error) {
	customer, err := repositories.FindCustomerById(id)
	if err != nil {
		return models.Customer{}, fmt.Errorf("error fetching customer with ID %s", id)
	}

	return customer, nil
}

func CreateCustomer(data map[string]interface{}) (models.Customer, error) {
	customer := models.Customer{
		ID: uuid.New(),
		FirstName: data["first_name"].(string),
		LastName: data["last_name"].(string),
		Nationality: data["nationality"].(string),
		IdentificationNumber: data["identification_number"].(string),
		IdentificationType: data["identification_type"].(string),
	}

	_, err := repositories.SaveCustomer(&customer)
	if err != nil {
		return models.Customer{}, fmt.Errorf("error saving customer, %v", err)
	}

	return customer, nil
}

func EditCustomer(data map[string]interface{}) (models.Customer, error) {
	customer, err := repositories.FindCustomerById(data["id"].(uuid.UUID))
	if err != nil {
		return models.Customer{}, fmt.Errorf("error fetching customer with ID %s", data["id"])
	}

	customer.Nationality = data["nationality"].(string)
	customer.FirstName = data["first_name"].(string)
	customer.LastName = data["last_name"].(string)
	customer.IdentificationNumber = data["identification_number"].(string)
	customer.IdentificationType = data["identification_type"].(string)

	_, err = repositories.UpdateCustomer(&customer)
	if err != nil {
		return models.Customer{}, fmt.Errorf("error occurred updating customer")
	}

	return customer, nil
}

func DeleteCustomer(id uuid.UUID) error {
	customer, err := repositories.FindCustomerById(id)
	if err != nil {
		return fmt.Errorf("error fetching customer with ID %s", id)
	}

	_, err = repositories.DeleteCustomer(customer.ID)
	if err != nil {
		return fmt.Errorf("error deleting customer with ID %s", id)
	}

	return nil
}
