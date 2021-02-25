package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "root"
	password string = "root"
	database string = "manajemen_mahasiswa"
)

var (
	dsn = fmt.Sprintf("%v:%v@tcp%v/%v", username, password, "(localhost:8889)", database)
)

// HubToMySQL
func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
