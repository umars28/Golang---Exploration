package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "root"
	password string = ""
	database string = "manajemen_mahasiswa"
	host     = "localhost"
	port     = 3306
)

var (
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
)

// HubToMySQL
func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
