package query

import (
	"database/sql"
	"fmt"
	"log"
	"mygo/config"
	"mygo/model"

	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func Register(name, nim, email, pass string) {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}
	password, _ := HashPassword(pass)
	_, error := db.Query(`INSERT INTO users (nim, nama, email, password) VALUES (?, ?, ?, ?)`, nim, name, email, password)

	// if there is an error inserting, handle it
	if error != nil {
		panic(error.Error())
	}
	fmt.Println("success")
}

func Login(email, password string) bool {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}
	//users := QueryUser(email)
	var tag model.User
	var isSuccess bool
	// Execute the query
	db.QueryRow("SELECT * FROM users where email = ?", email).Scan(&tag.ID, &tag.Nama, &tag.Nim, &tag.Email, &tag.Password)
	var password_tes = bcrypt.CompareHashAndPassword([]byte(tag.Password), []byte(password))
	if email == tag.Email && password_tes == nil {
		isSuccess = true
	} else {
		isSuccess = false
	}
	return isSuccess
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
