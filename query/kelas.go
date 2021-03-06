package query

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"mygo/config"
	"mygo/model"
)

func GetAllKelas(ctx context.Context) ([]model.Kelas, error) {
	var kelas []model.Kelas
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", "kelas")

	rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next() {
		var kls model.Kelas
		if err = rowQuery.Scan(&kls.ID,
			&kls.Nama); err != nil {
			return nil, err
		}

		kelas = append(kelas, kls)
	}
	return kelas, nil

}

func CreateRowKelas(name string) {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}
	_, error := db.Exec(`INSERT INTO kelas (nama) VALUES (?)`, name)
	if error != nil {
		panic(error.Error())
	}
	fmt.Println("success")
}

func KelasDelete(kelas int) {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", "kelas", kelas)
	_, error := db.Query(queryText)
	if error != nil {
		fmt.Println("failed")
	}

	if err != nil && err != sql.ErrNoRows {
		fmt.Println("gagal")
	}
	fmt.Println("successfully deleted")
	return
}

func KelasEdit(ctx context.Context, id int) ([]model.Kelas, error) {
	var kelas []model.Kelas
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	//queryText := fmt.Sprintf("SELECT * FROM %v where id %d Order By id DESC", table, mhs)
	rowQuery, err := db.Query("SELECT * FROM kelas where id = ?", id)

	//rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next() {
		var kls model.Kelas

		if err = rowQuery.Scan(&kls.ID,
			&kls.Nama); err != nil {
		}

		kelas = append(kelas, kls)
	}

	if err != nil {
		fmt.Println(err)
	}
	return kelas, nil
}

func KelasUpdate(id, name string) {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}
	_, error := db.Query(`UPDATE kelas SET nama = ? where id = ?`, name, id)
	if error != nil {
		panic(error.Error())
	}
}
