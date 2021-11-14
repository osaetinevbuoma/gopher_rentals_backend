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
