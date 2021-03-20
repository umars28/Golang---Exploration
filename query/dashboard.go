package query

import (
	"context"
	"fmt"
	"log"
	"mygo/config"
	"mygo/model"
)

func GetUser(ctx context.Context, email string) ([]model.User, error) {
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
	rowQuery, err := db.Query("SELECT * FROM users where email = ?", email)

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
