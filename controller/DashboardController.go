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

	_ "github.com/go-sql-driver/mysql"
)

func DashboardController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		email := r.URL.Query().Get("email")
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		user, err := query.GetUser(ctx, email)

		if err != nil {
			fmt.Println(err)
		}
		template, err := template.ParseFiles(
			path.Join("views/dashboard", "dashboard.html"),
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

		err = template.Execute(w, user)
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
