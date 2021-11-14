package repositories

import (
	"github.com/google/uuid"
	"gopher_rentals/db"
	"gopher_rentals/models"
	"gopher_rentals/repositories"
	"testing"
)

func TestSaveCustomer(t *testing.T) {
	if err := db.ConfigureDB(); err != nil {
		t.Fatalf("Connection established successfully")
	}

	customer := models.Customer{
		ID: uuid.New(),
		FirstName: "John",
		LastName: "Doe",
		Nationality: "Nigerian",
		IdentificationNumber: "ABC123",
		IdentificationType: "International Passport",
	}

	row, err := repositories.SaveCustomer(customer)
	if err != nil {
		t.Fatalf("TestSaveCustomer not saved: %v", err)
	}

	if row != 1 {
		t.Fatalf("TestSaveCustomer row not inserted")
	}
}

func TestFindAllCustomers(t *testing.T) {
	if err := db.ConfigureDB(); err != nil {
		t.Fatalf("Connection established successfully")
	}

	customers, err := repositories.FindAllCustomers()
	if err != nil {
		t.Fatalf("TestFindAllCustomers error: %v", err)
	}

	if len(customers) == 0 {
		t.Fatalf("TestFindAllCustomers length is 0")
	}
}

func TestFindCustomerUpdateCustomerDeleteCustomer(t *testing.T) {
	if err := db.ConfigureDB(); err != nil {
		t.Fatalf("Connection established successfully")
	}

	customers, err := repositories.FindAllCustomers()
	if err != nil {
		t.Fatalf("TestFindCustomerUpdateCustomerDeleteCustomer error: %v", err)
	}

	if len(customers) == 0 {
		t.Fatalf("TestFindCustomerUpdateCustomerDeleteCustomer length is 0")
	}

	customerId := customers[0].ID
	customer, err := repositories.FindCustomerById(customerId)
	if err != nil {
		t.Fatalf("TestFindCustomerUpdateCustomerDeleteCustomer error: %v", err)
	}

	if customer.ID != customerId {
		t.Fatalf("TestFindCustomerUpdateCustomerDeleteCustomer: Customer IDs do not match")
	}

	customer.FirstName = "John 2"
	customer.LastName = "Doe 2"

	row, err := repositories.UpdateCustomer(customer)
	if err != nil {
		t.Fatalf("TestFindCustomerUpdateCustomerDeleteCustomer error: %v", err)
	}

	if row != 1 {
		t.Fatalf("TestFindCustomerUpdateCustomerDeleteCustomer: update row is not 1")
	}

	row, err = repositories.DeleteCustomer(customer.ID)
	if err != nil {
		t.Fatalf("TestFindCustomerUpdateCustomerDeleteCustomer error: %v", err)
	}

	if row != 1 {
		t.Fatalf("TestFindCustomerUpdateCustomerDeleteCustomer: delete row is not 1")
	}
}
