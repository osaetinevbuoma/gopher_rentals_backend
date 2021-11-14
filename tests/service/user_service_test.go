package service

import (
	"gopher_rentals/db"
	"gopher_rentals/services"
	"testing"
)

func TestCreateUser(t *testing.T) {
	_ = db.ConfigureDB()

	user, err := services.CreateUser()
	if err != nil {
		t.Fatalf("TestCreateUser: %v", err)
	}

	isPasswordCorrect := services.CheckPassword(user.Password, "password")
	if !isPasswordCorrect {
		t.Fatalf("TestCreatUser: passwords do not match")
	}
}
