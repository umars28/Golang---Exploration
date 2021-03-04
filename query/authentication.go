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

func Register(name, nim, email, pass, role string) {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}
	password, _ := HashPassword(pass)
	res, error := db.Exec(`INSERT INTO users (nama, email, password, roles) VALUES (?, ?, ?, ?)`, name, email, password, role)
	// if there is an error inserting, handle it
	id, _ := res.LastInsertId()
	var semester int
	var kelas_id int
	_, error2 := db.Exec(`INSERT INTO mahasiswa (nim, name, semester , users_id, kelas_id) VALUES (?, ?, ?, ?, ?)`, nim, name, semester, id, kelas_id)
	if error != nil || error2 != nil {
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
	db.QueryRow("SELECT * FROM users where email = ?", email).Scan(&tag.ID, &tag.Nama, &tag.Email, &tag.Password, &tag.Roles)
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
