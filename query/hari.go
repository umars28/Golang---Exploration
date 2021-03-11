package query

import (
	"context"
	"fmt"
	"log"
	"mygo/config"
	"mygo/model"
)

func GetAllHari(ctx context.Context) ([]model.Hari, error) {
	var haris []model.Hari
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", "hari")

	rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next() {
		var hari model.Hari
		if err = rowQuery.Scan(&hari.ID,
			&hari.Nama); err != nil {
			return nil, err
		}

		haris = append(haris, hari)
	}
	return haris, nil

}
