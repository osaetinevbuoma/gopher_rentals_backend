package services

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopher_rentals/models"
	"gopher_rentals/repositories"
)

func CreateUser() (models.User, error) {
	password, err := HashPassword("password")
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		ID: uuid.New(),
		Email: "admin@email.com",
		Password: password,
	}
	
	_, err = repositories.SaveUser(&user)
	if err != nil {
		return models.User{}, fmt.Errorf("error occurred saving user -> %v", err)
	}
	
	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		return "", fmt.Errorf("failed to hash password, %s -> %v", password, err)
	}

	return string(bytes), nil
}

func CheckPassword(password string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(providedPassword))
	if err != nil {
		return false
	}

	return true
}
