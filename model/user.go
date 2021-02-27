package model

import "database/sql"

var db *sql.DB
var err error

type User struct {
	ID       int
	Nama     string
	Nim      string
	Email    string
	Password string
}
