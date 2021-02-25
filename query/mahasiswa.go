package query

import (
	"context"
	"fmt"
	"log"
	"mygo/config"
	"mygo/model"
	"time"
)

const (
	table          = "mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]model.Mahasiswa, error) {
	var mahasiswas []model.Mahasiswa
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next() {
		var mahasiswa model.Mahasiswa
		var created_at, updated_at string

		if err = rowQuery.Scan(&mahasiswa.ID,
			&mahasiswa.NIM,
			&mahasiswa.Name,
			&mahasiswa.Semester,
			&created_at,
			&updated_at); err != nil {
			return nil, err
		}

		// format string ke datetime
		mahasiswa.CreatedAt, err = time.Parse(layoutDateTime, created_at)
		if err != nil {
			log.Fatal(err)
		}
		mahasiswa.UpdatedAt, err = time.Parse(layoutDateTime, updated_at)
		if err != nil {
			log.Fatal(err)
		}

		mahasiswas = append(mahasiswas, mahasiswa)
	}
	return mahasiswas, nil

}

// func CreateRow(nim int, name string, semester int) {
// 	db, e := config.MySQL()

// 	if e != nil {
// 		log.Fatal("Can't connect to mysql", e)
// 	}

// 	eb := db.Ping()
// 	if eb != nil {
// 		panic(eb.Error())
// 	}

// 	data, error := db.Query(`INSERT INTO mahasiswa (nim, name, semester, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`, nim, name, semester, time.Now(), time.Now())

// 	// if there is an error inserting, handle it
// 	if error != nil {
// 		panic(error.Error())
// 	}
// 	fmt.Println("success")
// 	return data
// }

// func Delete(mhs model.Mahasiswa) {

// 	db, err := config.MySQL()

// 	if err != nil {
// 		log.Fatal("Can't connect to MySQL", err)
// 	}

// 	queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", table, mhs.ID)
// 	_, error := db.Query(queryText)
// 	if error != nil {
// 		fmt.Println("failed")
// 	}

// 	if err != nil && err != sql.ErrNoRows {
// 		fmt.Println("gagal")
// 	}
// 	fmt.Println("successfully deleted")
// 	return
// }
