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
