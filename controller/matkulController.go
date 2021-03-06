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

func MatkulController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		matkul, err := query.GetAllMatkul(ctx)

		if err != nil {
			fmt.Println(err)
		}
		template, err := template.ParseFiles(
			path.Join("views/matkul", "matkul.html"),
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

		err = template.Execute(w, matkul)
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

func TambahMatkulController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		template, err := template.ParseFiles(
			path.Join("views/matkul", "tambah.html"),
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

		err = template.Execute(w, nil)
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

func MatkulStoreController(w http.ResponseWriter, r *http.Request) {
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

	name := r.Form.Get("name")
	status := r.Form.Get("status")
	query.CreateRowMatkul(name, status)
	fmt.Println("success")
	http.Redirect(w, r, "/matkul", 302)
	w.Write([]byte("<script>alert('Sukses menambahkan data')</script>"))
	return
}

func MatkulDeleteController(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		config.MessageError500(w, r)
		fmt.Println(w, "id tidak boleh kosong", http.StatusBadRequest)
		return
	}
	matkul, _ := strconv.Atoi(id)
	query.MatkulDelete(matkul)
	fmt.Println("sukses hapus data")
	http.Redirect(w, r, "/matkul", 302)
	w.Write([]byte("<script>alert('Sukses menghapus data')</script>"))
	return
}

func MatkulEditController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		matkulId, _ := strconv.Atoi(id)
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		matkul, err := query.MatkulEdit(ctx, matkulId)
		template, err := template.ParseFiles(
			path.Join("views/matkul", "edit.html"),
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

		err = template.Execute(w, matkul)
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

func MatkulUpdateController(w http.ResponseWriter, r *http.Request) {
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
	name := r.Form.Get("name")
	status := r.Form.Get("status")
	query.MatkulUpdate(id, name, status)
	fmt.Println("success")
	http.Redirect(w, r, "/matkul", 302)
	w.Write([]byte("<script>alert('Sukses mengubah data')</script>"))
	return
}
