package query

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"mygo/config"
	"mygo/model"
)

func GetAllMatkul(ctx context.Context) ([]model.Matkul, error) {
	var matkuls []model.Matkul
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", "matkul")

	rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next() {
		var matkul model.Matkul
		if err = rowQuery.Scan(&matkul.ID,
			&matkul.Nama,
			&matkul.Status); err != nil {
			return nil, err
		}

		matkuls = append(matkuls, matkul)
	}
	return matkuls, nil

}

func CreateRowMatkul(name, status string) {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}
	_, error := db.Exec(`INSERT INTO matkul (name, status) VALUES (?, ?)`, name, status)
	if error != nil {
		panic(error.Error())
	}
	fmt.Println("success")
}

func MatkulDelete(kelas int) {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", "matkul", kelas)
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

func MatkulEdit(ctx context.Context, id int) ([]model.Matkul, error) {
	var matkuls []model.Matkul
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	//queryText := fmt.Sprintf("SELECT * FROM %v where id %d Order By id DESC", table, mhs)
	rowQuery, err := db.Query("SELECT * FROM matkul where id = ?", id)

	//rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next() {
		var matkul model.Matkul

		if err = rowQuery.Scan(&matkul.ID,
			&matkul.Nama,
			&matkul.Status); err != nil {
		}

		matkuls = append(matkuls, matkul)
	}

	if err != nil {
		fmt.Println(err)
	}
	return matkuls, nil
}

func MatkulUpdate(id, name, status string) {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}
	_, error := db.Query(`UPDATE matkul SET nama = ?, status = ? where id = ?`, name, status, id)
	if error != nil {
		panic(error.Error())
	}
}
