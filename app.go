package main

import (
	"log"
	"mygo/controller"
	"mygo/middleware"
	"net/http"
)

func main() {
	const port = ":8080"
	mux := http.NewServeMux()
	mux.Handle("/mahasiswa", middleware.Auth(http.HandlerFunc(controller.MahasiswaController)))
	mux.Handle("/tambah", middleware.Auth(http.HandlerFunc(controller.TambahController)))
	mux.Handle("/store", middleware.Auth(http.HandlerFunc(controller.StoreController)))
	mux.Handle("/delete", middleware.Auth(http.HandlerFunc(controller.DeleteController)))
	mux.Handle("/detail", middleware.Auth(http.HandlerFunc(controller.DetailController)))
	mux.Handle("/edit", middleware.Auth(http.HandlerFunc(controller.EditController)))
	mux.Handle("/update", middleware.Auth(http.HandlerFunc(controller.UpdateController)))
	mux.Handle("/register", middleware.CheckSession(http.HandlerFunc(controller.Register)))
	mux.Handle("/store/register", middleware.CheckSession(http.HandlerFunc(controller.StoreRegister)))
	mux.Handle("/login", middleware.CheckSession(http.HandlerFunc(controller.Login)))
	mux.Handle("/login/proses", middleware.CheckSession(http.HandlerFunc(controller.LoginProses)))
	mux.Handle("/logout", middleware.Auth(http.HandlerFunc(controller.Logout)))

	// mux.HandleFunc("/", controller.HomeController)

	fileServerCss := http.FileServer(http.Dir("assets/admin/css"))
	fileServerCssPlugins := http.FileServer(http.Dir("assets/admin/plugins/css"))
	fileServerJs := http.FileServer(http.Dir("assets/admin/js"))
	fileServerJsPlugins := http.FileServer(http.Dir("assets/admin/plugins/js"))
	fileServerImage := http.FileServer(http.Dir("assets/admin/image"))

	mux.Handle("/static/css/", http.StripPrefix("/static/css", fileServerCss))
	mux.Handle("/static/plugins/css/", http.StripPrefix("/static/plugins/css", fileServerCssPlugins))
	mux.Handle("/static/js/", http.StripPrefix("/static/js", fileServerJs))
	mux.Handle("/static/plugins/js/", http.StripPrefix("/static/plugins/js", fileServerJsPlugins))
	mux.Handle("/static/image/", http.StripPrefix("/static/image/", fileServerImage))

	log.Println("Starting web on port " + port)

	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
