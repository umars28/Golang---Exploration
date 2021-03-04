package query

import (
	"context"
	"fmt"
	"log"
	"mygo/config"
	"mygo/model"
)

func GetAllUser(ctx context.Context) ([]model.User, error) {
	var users []model.User
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table2)

	rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next() {
		var user model.User
		if err = rowQuery.Scan(&user.ID,
			&user.Nama,
			&user.Email,
			&user.Password,
			&user.Roles); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil

}

func EditUser(ctx context.Context, userId int) ([]model.User, error) {
	var users []model.User
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	//queryText := fmt.Sprintf("SELECT * FROM %v where id %d Order By id DESC", table, mhs)
	rowQuery, err := db.Query("SELECT * FROM users where id = ?", userId)
	//rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next() {
		var user model.User

		if err = rowQuery.Scan(&user.ID,
			&user.Nama,
			&user.Email,
			&user.Password,
			&user.Roles); err != nil {
		}

		users = append(users, user)
	}

	if err != nil {
		fmt.Println(err)
	}
	return users, nil
}

func UserUpdate(id, name, email, password, roles string) {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}
	if len(password) > 0 {
		pass, _ := HashPassword(password)
		_, error := db.Query(`UPDATE users SET nama = ?, email = ?, password = ? where id = ?`, name, email, pass, id)
		if error != nil {
			panic(error.Error())
		}
	}
	_, error2 := db.Query(`UPDATE users SET nama = ?, email = ? where id = ?`, name, email, id)
	if error2 != nil {
		panic(error2.Error())
	}
}
