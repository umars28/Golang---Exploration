package controller

import (
	"fmt"
	"html/template"
	"log"
	"mygo/query"
	"net/http"
	"path"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
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
		roles := "Mahasiswa"
		query.Register(name, nim, email, password, roles)
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
			session, _ := store.Get(r, "auth")
			session.Values["authenticated"] = true
			session.Save(r, w)
			fmt.Println("login berhasil")
			http.Redirect(w, r, "/mahasiswa", 302)
		} else {
			fmt.Println("login gagal")
			http.Redirect(w, r, "/login", 302)
		}
		return
	}
}

func Logout(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "auth")
	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(req, res)
	http.Redirect(res, req, "/login", 302)
}
