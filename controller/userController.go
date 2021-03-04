package controller

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"mygo/query"
	"net/http"
	"path"
	"strconv"
)

func UserController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		user, err := query.GetAllUser(ctx)

		if err != nil {
			fmt.Println(err)
		}
		template, err := template.ParseFiles(
			path.Join("views/user", "user.html"),
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

		err = template.Execute(w, user)
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

func UserEditController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		userId, _ := strconv.Atoi(id)
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		users, err := query.EditUser(ctx, userId)
		template, err := template.ParseFiles(
			path.Join("views/user", "edit.html"),
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

		err = template.Execute(w, users)
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

func UserUpdateController(w http.ResponseWriter, r *http.Request) {
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
	nama := r.Form.Get("nama")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	roles := r.Form.Get("roles")
	query.UserUpdate(id, nama, email, password, roles)
	fmt.Println("success")
	http.Redirect(w, r, "/user", 302)
	w.Write([]byte("<script>alert('Sukses mengubah data')</script>"))
	return
}
