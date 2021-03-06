package query

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"mygo/config"
	"mygo/model"
	"strconv"
)

func GetAllDosen(ctx context.Context) ([]model.Dosen, error) {
	var dosens []model.Dosen
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", "dosen")

	rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next() {
		var dosen model.Dosen
		if err = rowQuery.Scan(&dosen.ID,
			&dosen.NIP,
			&dosen.Name,
			&dosen.Email,
			&dosen.Matkul_id,
			&dosen.User_id); err != nil {
			return nil, err
		}

		dosens = append(dosens, dosen)
	}
	return dosens, nil

}

func CreateRowDosen(name, nip, email, matkul_id string) {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}
	role := "Dosen"
	password, _ := HashPassword(nip)
	res, error := db.Exec(`INSERT INTO users (nama, email, password, roles) VALUES (?, ?, ?, ?)`, name, email, password, role)
	id, _ := res.LastInsertId()
	fmt.Println(id)
	_, error2 := db.Exec(`INSERT INTO dosen (nip, nama, email, matkul_id, users_id) VALUES (?, ?, ?, ?, ?)`, nip, name, email, matkul_id, id)
	if error != nil || error2 != nil {
		panic(error.Error())
	}
	fmt.Println("success")
}

func DeleteDosen(dosen, user_id int) {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", "dosen", dosen)
	queryText2 := fmt.Sprintf("DELETE FROM %v where id = '%d'", "users", user_id)
	_, error := db.Query(queryText)
	_, error2 := db.Query(queryText2)
	if error != nil || error2 != nil {
		fmt.Println("failed")
	}

	if err != nil && err != sql.ErrNoRows {
		fmt.Println("gagal")
	}
	fmt.Println("successfully deleted")
	return
}

func DetailDosen(ctx context.Context, dosenId int) ([]model.Dosen, error) {
	var dosens []model.Dosen
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	//queryText := fmt.Sprintf("SELECT * FROM %v where id %d Order By id DESC", table, mhs)
	rowQuery, err := db.Query("SELECT * FROM dosen where id = ?", dosenId)

	//rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next() {
		var dosen model.Dosen

		if err = rowQuery.Scan(&dosen.ID,
			&dosen.NIP,
			&dosen.Name,
			&dosen.Email,
			&dosen.Matkul_id,
			&dosen.User_id); err != nil {
		}

		dosens = append(dosens, dosen)
	}

	if err != nil {
		fmt.Println(err)
	}
	return dosens, nil
}

func EditDosen(ctx context.Context, dosenId int) ([]model.Dosen, error) {
	var dosens []model.Dosen
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	//queryText := fmt.Sprintf("SELECT * FROM %v where id %d Order By id DESC", table, mhs)
	rowQuery, err := db.Query("SELECT * FROM dosen where id = ?", dosenId)

	//rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next() {
		var dosen model.Dosen

		if err = rowQuery.Scan(&dosen.ID,
			&dosen.NIP,
			&dosen.Name,
			&dosen.Email,
			&dosen.Matkul_id,
			&dosen.User_id); err != nil {
		}

		dosens = append(dosens, dosen)
	}

	if err != nil {
		fmt.Println(err)
	}
	return dosens, nil
}

func UpdateDosen(id, nip, name, email, matkul_id, userId string) {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}
	password, _ := HashPassword(nip)
	user, _ := strconv.Atoi(userId)
	_, error := db.Query(`UPDATE dosen SET nip = ?, nama = ?, email = ?, matkul_id = ?, kelas_id = ? where id = ?`, nip, name, email, matkul_id, id)
	_, error2 := db.Query(`UPDATE users SET nama = ?, password = ? where id = ?`, name, password, user)
	if error != nil || error2 != nil {
		panic(error.Error())
	}
}
