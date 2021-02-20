package main

import (
	"log"
	"mygo/controller"
	"net/http"
)

func main() {
	const port = ":8080"
	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.HomeController)

	fileServerCss := http.FileServer(http.Dir("assets/admin/css"))
	fileServerCssPlugins := http.FileServer(http.Dir("assets/admin/plugins/css"))
	fileServerJs := http.FileServer(http.Dir("assets/admin/js"))
	fileServerJsPlugins := http.FileServer(http.Dir("assets/admin/plugins/js"))

	mux.Handle("/static/css/", http.StripPrefix("/static/css", fileServerCss))
	mux.Handle("/static/plugins/css/", http.StripPrefix("/static/plugins/css", fileServerCssPlugins))
	mux.Handle("/static/js/", http.StripPrefix("/static/js", fileServerJs))
	mux.Handle("/static/plugins/js/", http.StripPrefix("/static/plugins/js", fileServerJsPlugins))

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
