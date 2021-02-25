package controller

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	template, err := template.ParseFiles(
		path.Join("views", "home.html"),
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

}
