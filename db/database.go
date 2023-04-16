package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strings"
)

var DB *sql.DB

func ConfigureDB() error {
	config := mysql.Config{
		User:   os.Getenv("MARIADB_USER"),
		Passwd: os.Getenv("MARIADB_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("MARIADB_HOSTNAME") + ":" + os.Getenv("MARIADB_PORT"),
		//DBName:               os.Getenv("MARIADB_DATABASE"),
		AllowNativePasswords: true,
		ParseTime:            true,
		MultiStatements:      true,
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

	dbQueries, err := os.ReadFile("db/database.sql")
	if err != nil {
		log.Fatal("DB sql file not provided")
		return err
	}

	if _, err := DB.Exec(strings.TrimSpace(string(dbQueries))); err != nil {
		log.Fatal("Failed to execute DB queries", err)
		return err
	}

	log.Println("Database connection successful")
	return nil
}
