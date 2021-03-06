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

func KelasController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		kelas, err := query.GetAllKelas(ctx)

		if err != nil {
			fmt.Println(err)
		}
		template, err := template.ParseFiles(
			path.Join("views/kelas", "kelas.html"),
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

		err = template.Execute(w, kelas)
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

func TambahKelasController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		template, err := template.ParseFiles(
			path.Join("views/kelas", "tambah.html"),
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

func KelasStoreController(w http.ResponseWriter, r *http.Request) {
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
	query.CreateRowKelas(name)
	fmt.Println("success")
	http.Redirect(w, r, "/kelas", 302)
	w.Write([]byte("<script>alert('Sukses menambahkan data')</script>"))
	return
}

func KelasDeleteController(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		config.MessageError500(w, r)
		fmt.Println(w, "id tidak boleh kosong", http.StatusBadRequest)
		return
	}
	kelas, _ := strconv.Atoi(id)
	query.KelasDelete(kelas)
	fmt.Println("sukses hapus data")
	http.Redirect(w, r, "/kelas", 302)
	w.Write([]byte("<script>alert('Sukses menghapus data')</script>"))
	return
}

func KelasEditController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		kelas, _ := strconv.Atoi(id)
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		kls, err := query.KelasEdit(ctx, kelas)
		template, err := template.ParseFiles(
			path.Join("views/kelas", "edit.html"),
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

		err = template.Execute(w, kls)
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

func KelasUpdateController(w http.ResponseWriter, r *http.Request) {
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
	query.KelasUpdate(id, name)
	fmt.Println("success")
	http.Redirect(w, r, "/kelas", 302)
	w.Write([]byte("<script>alert('Sukses mengubah data')</script>"))
	return
}
