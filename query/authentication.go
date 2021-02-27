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

func QueryUser(email string) model.User {
	var users = model.User{}
	db.QueryRow(`
		SELECT id, 
		nama, 
		nim, 
		email, 
		password 
		FROM users WHERE email=?
		`, email).
		Scan(
			&users.ID,
			&users.Nama,
			&users.Nim,
			&users.Email,
			&users.Password,
		)
	return users
}

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

func Login(email, password string) {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}
	users := QueryUser(email)
	//deskripsi dan compare password
	// var password_tes = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password)

	// if password_tes == nil {
	// 	//login success
	// 	fmt.Println("sukses login bree")
	// } else {
	// 	//login failed
	// 	fmt.Println("gagal login bro")
	// }
	CheckPasswordHash(users.Password, password)
	fmt.Println("success")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
