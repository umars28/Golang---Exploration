package query

import (
	"context"
	"fmt"
	"log"
	"mygo/config"
	"mygo/model"
)

func GetAllJadwal(ctx context.Context) ([]model.Jadwal, error) {
	var jadwals []model.Jadwal
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", "jadwal")

	rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next() {
		var jadwal model.Jadwal
		if err = rowQuery.Scan(&jadwal.ID,
			&jadwal.Hari_id,
			&jadwal.Dosen_id,
			&jadwal.Kelas_id,
			&jadwal.Matkul_id); err != nil {
			return nil, err
		}

		jadwals = append(jadwals, jadwal)
	}
	return jadwals, nil

}

// func CreateRow(name, nim, semester, email, kelas_id string) {
// 	db, e := config.MySQL()

// 	if e != nil {
// 		log.Fatal("Can't connect to mysql", e)
// 	}

// 	eb := db.Ping()
// 	if eb != nil {
// 		panic(eb.Error())
// 	}
// 	role := "Mahasiswa"
// 	password, _ := HashPassword(nim)
// 	res, error := db.Exec(`INSERT INTO users (nama, email, password, roles) VALUES (?, ?, ?, ?)`, name, email, password, role)
// 	id, _ := res.LastInsertId()
// 	fmt.Println(id)
// 	_, error2 := db.Exec(`INSERT INTO mahasiswa (nim, name, semester , users_id, kelas_id) VALUES (?, ?, ?, ?, ?)`, nim, name, semester, id, kelas_id)
// 	if error != nil || error2 != nil {
// 		panic(error.Error())
// 	}
// 	fmt.Println("success")
// }

// func Delete(mhs, user_id int) {

// 	db, err := config.MySQL()

// 	if err != nil {
// 		log.Fatal("Can't connect to MySQL", err)
// 	}

// 	queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", table, mhs)
// 	queryText2 := fmt.Sprintf("DELETE FROM %v where id = '%d'", table2, user_id)
// 	_, error := db.Query(queryText)
// 	_, error2 := db.Query(queryText2)
// 	if error != nil || error2 != nil {
// 		fmt.Println("failed")
// 	}

// 	if err != nil && err != sql.ErrNoRows {
// 		fmt.Println("gagal")
// 	}
// 	fmt.Println("successfully deleted")
// 	return
// }

// func Detail(ctx context.Context, mhs int) ([]model.Mahasiswa, error) {
// 	var mahasiswas []model.Mahasiswa
// 	db, e := config.MySQL()

// 	if e != nil {
// 		log.Fatal("Can't connect to mysql", e)
// 	}

// 	eb := db.Ping()
// 	if eb != nil {
// 		panic(eb.Error())
// 	}

// 	//queryText := fmt.Sprintf("SELECT * FROM %v where id %d Order By id DESC", table, mhs)
// 	rowQuery, err := db.Query("SELECT * FROM mahasiswa where id = ?", mhs)

// 	//rowQuery, err := db.QueryContext(ctx, queryText)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for rowQuery.Next() {
// 		var mahasiswa model.Mahasiswa

// 		if err = rowQuery.Scan(&mahasiswa.ID,
// 			&mahasiswa.NIM,
// 			&mahasiswa.Name,
// 			&mahasiswa.Semester,
// 			&mahasiswa.User_id,
// 			&mahasiswa.Kelas_id); err != nil {
// 		}

// 		mahasiswas = append(mahasiswas, mahasiswa)
// 	}

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return mahasiswas, nil
// }

// func Edit(ctx context.Context, mhs int) ([]model.Mahasiswa, error) {
// 	var mahasiswas []model.Mahasiswa
// 	db, e := config.MySQL()

// 	if e != nil {
// 		log.Fatal("Can't connect to mysql", e)
// 	}

// 	eb := db.Ping()
// 	if eb != nil {
// 		panic(eb.Error())
// 	}

// 	//queryText := fmt.Sprintf("SELECT * FROM %v where id %d Order By id DESC", table, mhs)
// 	rowQuery, err := db.Query("SELECT * FROM mahasiswa where id = ?", mhs)

// 	//rowQuery, err := db.QueryContext(ctx, queryText)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for rowQuery.Next() {
// 		var mahasiswa model.Mahasiswa

// 		if err = rowQuery.Scan(&mahasiswa.ID,
// 			&mahasiswa.NIM,
// 			&mahasiswa.Name,
// 			&mahasiswa.Semester,
// 			&mahasiswa.User_id,
// 			&mahasiswa.Kelas_id); err != nil {
// 		}

// 		mahasiswas = append(mahasiswas, mahasiswa)
// 	}

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return mahasiswas, nil
// }

// func Update(id, nim, name, semester, kelas, userId string) {
// 	db, e := config.MySQL()

// 	if e != nil {
// 		log.Fatal("Can't connect to mysql", e)
// 	}

// 	eb := db.Ping()
// 	if eb != nil {
// 		panic(eb.Error())
// 	}
// 	password, _ := HashPassword(nim)
// 	user, _ := strconv.Atoi(userId)
// 	_, error := db.Query(`UPDATE mahasiswa SET nim = ?, name = ?, semester = ?, kelas_id = ? where id = ?`, nim, name, semester, kelas, id)
// 	_, error2 := db.Query(`UPDATE users SET nama = ?, password = ? where id = ?`, name, password, user)
// 	if error != nil || error2 != nil {
// 		panic(error.Error())
// 	}
// }
