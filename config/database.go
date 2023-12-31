package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnect() (*sql.DB, error) {
	user := getEnvValue("MYSQL_USER")
	password := getEnvValue("MYSQL_PASSWORD")
	host := getEnvValue("MYSQL_HOST")
	port := getEnvValue("MYSQL_PORT")
	database := getEnvValue("MYSQL_DATABASE")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
