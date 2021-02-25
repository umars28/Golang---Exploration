package controller

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"mygo/config"
	"mygo/model"
	"mygo/query"
	"net/http"
	"path"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MahasiswaController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		mahasiswas, err := query.GetAll(ctx)

		if err != nil {
			fmt.Println(err)
		}
		template, err := template.ParseFiles(
			path.Join("views/mahasiswa", "mahasiswa.html"),
			path.Join("views/template", "main.html"),
			path.Join("views/template", "header.html"),
			path.Join("views/template", "sidebar.html"),
			path.Join("views/template", "footer.html"),
		)
		if err != nil {
			log.Println(err)
			http.Error(w, "error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		err = template.Execute(w, mahasiswas)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calms", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}

func TambahController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		template, err := template.ParseFiles(
			path.Join("views/mahasiswa", "tambah.html"),
			path.Join("views/template", "main.html"),
			path.Join("views/template", "header.html"),
			path.Join("views/template", "sidebar.html"),
			path.Join("views/template", "footer.html"),
		)
		if err != nil {
			log.Println(err)
			http.Error(w, "error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		err = template.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calms", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}

var db *sql.DB

func StoreController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "Erorr is happening, keep calms", http.StatusInternalServerError)
		return
	}

	nim := r.Form.Get("nim")
	name := r.Form.Get("name")
	semester := r.Form.Get("semester")
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	_, error := db.Query(`INSERT INTO mahasiswa (nim, name, semester, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`, nim, name, semester, time.Now(), time.Now())

	// if there is an error inserting, handle it
	if error != nil {
		panic(error.Error())
	}
	fmt.Println("success")
	w.Write([]byte("Sukses Tambah Data"))
}

func DeleteController(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		fmt.Println(w, "id tidak boleh kosong", http.StatusBadRequest)
		return
	}
	mhs, _ := strconv.Atoi(id)

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", "mahasiswa", mhs)
	_, error := db.Query(queryText)
	if error != nil {
		fmt.Println("failed")
	}

	if err != nil && err != sql.ErrNoRows {
		fmt.Println("gagal")
	}
	w.Write([]byte("Sukses Hapus Data"))
	fmt.Println("sukses hapus data")
}

const (
	table          = "mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

func DetailController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		mhs, _ := strconv.Atoi(id)
		//ctx, cancel := context.WithCancel(context.Background())

		//defer cancel()
		//mahasiswas, err := query.GetAll(ctx)
		var mahasiswas []model.Mahasiswa
		db, e := config.MySQL()

		if e != nil {
			log.Fatal("Can't connect to mysql", e)
		}

		eb := db.Ping()
		if eb != nil {
			panic(eb.Error())
		}

		//queryText := fmt.Sprintf("SELECT * FROM %v where id %d Order By id DESC", table, mhs)
		rowQuery, err := db.Query("SELECT * FROM mahasiswa where id = ?", mhs)

		//rowQuery, err := db.QueryContext(ctx, queryText)
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

		if err != nil {
			fmt.Println(err)
		}
		template, err := template.ParseFiles(
			path.Join("views/mahasiswa", "detail.html"),
			path.Join("views/template", "main.html"),
			path.Join("views/template", "header.html"),
			path.Join("views/template", "sidebar.html"),
			path.Join("views/template", "footer.html"),
		)
		if err != nil {
			log.Println(err)
			http.Error(w, "error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		err = template.Execute(w, mahasiswas)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calms", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}

func EditController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		mhs, _ := strconv.Atoi(id)
		//ctx, cancel := context.WithCancel(context.Background())

		//defer cancel()
		//mahasiswas, err := query.GetAll(ctx)
		var mahasiswas []model.Mahasiswa
		db, e := config.MySQL()

		if e != nil {
			log.Fatal("Can't connect to mysql", e)
		}

		eb := db.Ping()
		if eb != nil {
			panic(eb.Error())
		}

		//queryText := fmt.Sprintf("SELECT * FROM %v where id %d Order By id DESC", table, mhs)
		rowQuery, err := db.Query("SELECT * FROM mahasiswa where id = ?", mhs)

		//rowQuery, err := db.QueryContext(ctx, queryText)
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

		if err != nil {
			fmt.Println(err)
		}
		template, err := template.ParseFiles(
			path.Join("views/mahasiswa", "edit.html"),
			path.Join("views/template", "main.html"),
			path.Join("views/template", "header.html"),
			path.Join("views/template", "sidebar.html"),
			path.Join("views/template", "footer.html"),
		)
		if err != nil {
			log.Println(err)
			http.Error(w, "error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		err = template.Execute(w, mahasiswas)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calms", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}

func UpdateController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "Erorr is happening, keep calms", http.StatusInternalServerError)
		return
	}

	id := r.Form.Get("id")
	nim := r.Form.Get("nim")
	name := r.Form.Get("name")
	semester := r.Form.Get("semester")
	db, e := config.MySQL()

	if e != nil {
		log.Fatal("Can't connect to mysql", e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	_, error := db.Query(`UPDATE mahasiswa SET nim = ?, name = ?, semester = ?, created_at = ?, updated_at = ? where id = ?`, nim, name, semester, time.Now(), time.Now(), id)
	if error != nil {
		panic(error.Error())
	}
	fmt.Println("success")
	w.Write([]byte("Sukses Update Data"))
}
