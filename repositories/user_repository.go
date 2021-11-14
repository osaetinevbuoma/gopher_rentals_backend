package repositories

import (
	"database/sql"
	"fmt"
	"gopher_rentals/db"
	"gopher_rentals/models"
)

func FindUserByEmail(email string) (models.User, error) {
	var user models.User

	row := db.DB.QueryRow("SELECT * FROM users WHERE email = ?", email)
	if err := row.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("FindUserByEmail: No user with email %s", email)
		}

		return user, fmt.Errorf("FindUserByEmail: %v", err)
	}

	return user, nil
}

func SaveUser(user *models.User) (int64, error) {
	result, err := db.DB.Exec("INSERT INTO users (id, email, password) VALUES (?, ?, ?)",
		user.ID, user.Email, user.Password)
	if err != nil {
		return 0, fmt.Errorf("SaveUser: Error occurred saving user -> %v", err)
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("SaveUser: No rows affected -> %v", err)
	}

	return row, nil
}
