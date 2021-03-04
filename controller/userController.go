package controller

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"mygo/query"
	"net/http"
	"path"
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

func TambahUserController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		template, err := template.ParseFiles(
			path.Join("views/user", "tambah.html"),
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
