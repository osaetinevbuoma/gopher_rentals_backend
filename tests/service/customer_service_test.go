package service

import (
	"github.com/google/uuid"
	"gopher_rentals/db"
	"gopher_rentals/services"
	"testing"
)

func TestCreateService(t *testing.T) {
	_ = db.ConfigureDB()

	customer1 := map[string]interface{}{
		"first_name": "John",
		"last_name": "Doe",
		"nationality": "Nigerian",
		"identification_number": "ABC123",
		"identification_type": "International Passport",
	}

	customer2 := map[string]interface{}{
		"first_name": "Jane",
		"last_name": "Doe",
		"nationality": "Australian",
		"identification_number": "XYZ123",
		"identification_type": "International Passport",
	}

	_, err := services.CreateCustomer(customer1)
	if err != nil {
		t.Fatalf("TestCreateService: %v", err)
	}

	_, err = services.CreateCustomer(customer2)
}

func TestListCustomers(t *testing.T) {
	_ = db.ConfigureDB()

	customers, err := services.ListCustomers()
	if err != nil {
		t.Fatalf("TestListCustomers: error occurred listing customers -> %v", err)
	}

	if len(customers) != 2 {
		t.Fatalf("TestListCustomers: customer length != 2")
	}
}

func TestGetCustomer(t *testing.T) {
	_ = db.ConfigureDB()

	customers, err := services.ListCustomers()
	if err != nil {
		t.Fatalf("TestGetCustomer: error occurred listing customers -> %v", err)
	}

	customer := customers[0]

	c, err := services.GetCustomer(customer.ID)
	if err != nil {
		t.Fatalf("TestGetCustomer: error getting customer -> %v", err)
	}

	if c.ID != customer.ID {
		t.Fatalf("TestGetCustomer: customer ID = %s.!= c ID %s", customer.ID, c.ID)
	}
}

func TestEditCustomer(t *testing.T) {
	_ = db.ConfigureDB()

	customers, err := services.ListCustomers()
	if err != nil {
		t.Fatalf("TestEditCustomer: error occurred listing customers -> %v", err)
	}

	customer := customers[0]
	updatedCustomer := map[string]interface{}{
		"id": customer.ID,
		"first_name": "Jonathan",
		"last_name": "Doe",
		"nationality": "Nigerian",
		"identification_number": "ABC123",
		"identification_type": "International Passport",
	}

	c, err := services.EditCustomer(updatedCustomer)
	if err != nil {
		t.Fatalf("TestEditCustomer: error occurred editing customer -> %v", err)
	}

	if c.ID != updatedCustomer["id"].(uuid.UUID) && c.FirstName != updatedCustomer["first_name"].(string) {
		t.Fatalf("TestEditCustomer: update customer does not match saved customer")
	}
}

func TestDeleteCustomer(t *testing.T) {
	_ = db.ConfigureDB()

	customers, err := services.ListCustomers()
	if err != nil {
		t.Fatalf("TestEditCustomer: error occurred listing customers -> %v", err)
	}

	for _, customer := range customers {
		err := services.DeleteCustomer(customer.ID)
		if err != nil {
			t.Fatalf("TestEditCutomer: error deleting customer -> %v", err)
		}
	}
}
