package db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func ConfigureDB() error {
	config := mysql.Config{
		User: "admin",
		Passwd: "",
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "gopher_rentals",
		AllowNativePasswords: true,
		ParseTime: true,
	}

	var err error
	DB, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
		return err
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		return err
	}

	fmt.Println("Database connection successful")
	return nil
}
