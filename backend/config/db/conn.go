package config

import (
	"database/sql"
	"log"
	"os"

	"backend/internal/helpers"

	"github.com/go-sql-driver/mysql"
)

func ConnMySql() (*sql.DB, error) {
	helpers.LoadEnvironmentFile()

	db_owner := os.Getenv("MYSQL_DB_OWNER")
	db_paswd := os.Getenv("MYSQL_DB_PASSWD")
	db_name := os.Getenv("MYSQL_DB_NAME")

	config := mysql.Config{
		DBName: db_name,
		User:   db_owner,
		Passwd: db_paswd,
		Net:    "tcp",
		Addr:   "localhost:3306",
		Params: map[string]string{
			"charset":         "utf8mb4",
			"collation":       "utf8mb4_unicode_ci",
			"parseTime":       "true",
			"multiStatements": "true",
		},
	}

	conn, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatalln("Unable to connect to database")
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		log.Fatalln("Failed while pinging database")
		return nil, err
	}

	log.Println("Database connection established")
	return conn, nil
}
