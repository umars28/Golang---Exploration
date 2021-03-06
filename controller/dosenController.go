package controller

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"mygo/config"
	"mygo/query"
	"net/http"
	"path"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func DosenController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		listDosen, err := query.GetAllDosen(ctx)

		if err != nil {
			fmt.Println(err)
		}
		template, err := template.ParseFiles(
			path.Join("views/dosen", "dosen.html"),
			path.Join("views/template", "main.html"),
			path.Join("views/template", "header.html"),
			path.Join("views/template", "sidebar.html"),
			path.Join("views/template", "footer.html"),
		)
		if err != nil {
			log.Println(err)
			config.MessageError500(w, r)
			return
		}

		err = template.Execute(w, listDosen)
		if err != nil {
			log.Println(err)
			config.MessageError500(w, r)
			return
		}
		return
	}

	config.MessageError503(w, r)
	return
}

func DosenTambahController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		listMatkul, err := query.GetAllMatkul(ctx)
		if err != nil {
			config.MessageError503(w, r)
			return
		}
		template, err := template.ParseFiles(
			path.Join("views/dosen", "tambah.html"),
			path.Join("views/template", "main.html"),
			path.Join("views/template", "header.html"),
			path.Join("views/template", "sidebar.html"),
			path.Join("views/template", "footer.html"),
		)
		if err != nil {
			log.Println(err)
			config.MessageError500(w, r)
			return
		}

		err = template.Execute(w, listMatkul)
		if err != nil {
			log.Println(err)
			config.MessageError500(w, r)
			return
		}
		return
	}

	config.MessageError503(w, r)
	return
}

func DosenStoreController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		config.MessageError500(w, r)
		return
	}

	nip := r.Form.Get("nip")
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	matkul_id := r.Form.Get("matkul_id")
	query.CreateRowDosen(name, nip, email, matkul_id)
	fmt.Println("success")
	http.Redirect(w, r, "/dosen", 302)
	w.Write([]byte("<script>alert('Sukses menambahkan data')</script>"))
	return
}

func DosenDeleteController(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	user_id := r.URL.Query().Get("userId")
	if id == "" {
		config.MessageError500(w, r)
		fmt.Println(w, "id tidak boleh kosong", http.StatusBadRequest)
		return
	}
	dosen, _ := strconv.Atoi(id)
	user, _ := strconv.Atoi(user_id)
	query.DeleteDosen(dosen, user)
	fmt.Println("sukses hapus data")
	http.Redirect(w, r, "/dosen", 302)
	w.Write([]byte("<script>alert('Sukses menghapus data')</script>"))
	return
}

func DosenDetailController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		dosen, _ := strconv.Atoi(id)
		listDosen, err := query.DetailDosen(ctx, dosen)
		if err != nil {
			log.Println(err)
			config.MessageError500(w, r)
			return
		}
		template, err := template.ParseFiles(
			path.Join("views/dosen", "detail.html"),
			path.Join("views/template", "main.html"),
			path.Join("views/template", "header.html"),
			path.Join("views/template", "sidebar.html"),
			path.Join("views/template", "footer.html"),
		)
		if err != nil {
			log.Println(err)
			config.MessageError500(w, r)
			return
		}

		err = template.Execute(w, listDosen)
		if err != nil {
			log.Println(err)
			config.MessageError500(w, r)
			return
		}
		return
	}

	config.MessageError503(w, r)
	return
}

func DosenEditController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		dosen, _ := strconv.Atoi(id)
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		listDosen, err := query.EditDosen(ctx, dosen)
		template, err := template.ParseFiles(
			path.Join("views/dosen", "edit.html"),
			path.Join("views/template", "main.html"),
			path.Join("views/template", "header.html"),
			path.Join("views/template", "sidebar.html"),
			path.Join("views/template", "footer.html"),
		)
		if err != nil {
			log.Println(err)
			config.MessageError500(w, r)
			return
		}

		err = template.Execute(w, listDosen)
		if err != nil {
			log.Println(err)
			config.MessageError500(w, r)
			return
		}
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}

func DosenUpdateController(w http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Query().Get("userId")
	if r.Method != http.MethodPost {
		config.MessageError503(w, r)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		config.MessageError500(w, r)
		return
	}

	id := r.Form.Get("id")
	nim := r.Form.Get("nip")
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	matkul_id := r.Form.Get("matkul_id")
	query.UpdateDosen(id, nim, name, email, matkul_id, user_id)
	fmt.Println("success")
	http.Redirect(w, r, "/dosen", 302)
	w.Write([]byte("<script>alert('Sukses mengubah data')</script>"))
	return
}
