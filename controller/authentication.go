package controller

import (
	"fmt"
	"html/template"
	"log"
	"mygo/query"
	"net/http"
	"path"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
		return
	}
	template, err := template.ParseFiles(
		path.Join("views/auth", "register.html"),
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

func StoreRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Erorr is happening, keep calms", http.StatusInternalServerError)
			return
		}
		password := r.Form.Get("password")
		password2 := r.Form.Get("password-confirm")
		if password != password2 {
			http.Error(w, "Mohon maaf, Konfirmasi Password Harus Sama", http.StatusInternalServerError)
			return
		}
		name := r.Form.Get("name")
		nim := r.Form.Get("nim")
		email := r.Form.Get("email")
		query.Register(name, nim, email, password)
		fmt.Println("success")
		http.Redirect(w, r, "/login", 302)
		w.Write([]byte("<script>alert('Sukses menambahkan data')</script>"))
		return

	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
		return
	}
	template, err := template.ParseFiles(
		path.Join("views/auth", "login.html"),
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

func LoginProses(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Erorr is happening, keep calms", http.StatusInternalServerError)
			return
		}
		password := r.Form.Get("password")
		email := r.Form.Get("email")
		success := query.Login(email, password)
		if success {
			fmt.Println("login berhasil")
			http.Redirect(w, r, "/mahasiswa", 302)
		} else {
			fmt.Println("login gagal")
			http.Redirect(w, r, "/login", 302)
		}
		return
	}
}
