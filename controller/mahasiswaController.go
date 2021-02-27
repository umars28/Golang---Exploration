package controller

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"mygo/query"
	"net/http"
	"path"
	"strconv"

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
	query.CreateRow(name, nim, semester)
	fmt.Println("success")
	http.Redirect(w, r, "/mahasiswa", 302)
	w.Write([]byte("<script>alert('Sukses menambahkan data')</script>"))
	return
}

func DeleteController(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		fmt.Println(w, "id tidak boleh kosong", http.StatusBadRequest)
		return
	}
	mhs, _ := strconv.Atoi(id)
	query.Delete(mhs)
	fmt.Println("sukses hapus data")
	http.Redirect(w, r, "/mahasiswa", 302)
	w.Write([]byte("<script>alert('Sukses menghapus data')</script>"))
	return
}

const (
	table          = "mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

func DetailController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		mhs, _ := strconv.Atoi(id)
		mahasiswas, err := query.Detail(ctx, mhs)
		if err != nil {
			log.Println(err)
			http.Error(w, "error is happening, keep calm", http.StatusInternalServerError)
			return
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
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		mahasiswas, err := query.Edit(ctx, mhs)
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
	query.Update(id, nim, name, semester)
	fmt.Println("success")
	http.Redirect(w, r, "/mahasiswa", 302)
	w.Write([]byte("<script>alert('Sukses mengubah data')</script>"))
	return
}