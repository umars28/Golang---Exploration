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
	// Mahasiswa
	mux.Handle("/mahasiswa", middleware.Auth(http.HandlerFunc(controller.MahasiswaController)))
	mux.Handle("/tambah", middleware.Auth(http.HandlerFunc(controller.TambahController)))
	mux.Handle("/store", middleware.Auth(http.HandlerFunc(controller.StoreController)))
	mux.Handle("/delete", middleware.Auth(http.HandlerFunc(controller.DeleteController)))
	mux.Handle("/detail", middleware.Auth(http.HandlerFunc(controller.DetailController)))
	mux.Handle("/edit", middleware.Auth(http.HandlerFunc(controller.EditController)))
	mux.Handle("/update", middleware.Auth(http.HandlerFunc(controller.UpdateController)))
	// End Mahasiswa

	// Auth
	mux.Handle("/register", middleware.CheckSession(http.HandlerFunc(controller.Register)))
	mux.Handle("/store/register", middleware.CheckSession(http.HandlerFunc(controller.StoreRegister)))
	mux.Handle("/login", middleware.CheckSession(http.HandlerFunc(controller.Login)))
	mux.Handle("/login/proses", middleware.CheckSession(http.HandlerFunc(controller.LoginProses)))
	mux.Handle("/logout", middleware.Auth(http.HandlerFunc(controller.Logout)))
	// End Auth

	// User
	mux.Handle("/user", middleware.Auth(http.HandlerFunc(controller.UserController)))
	mux.Handle("/edit-user", middleware.Auth(http.HandlerFunc(controller.UserEditController)))
	mux.Handle("/update-user", middleware.Auth(http.HandlerFunc(controller.UserUpdateController)))
	// End User

	// Kelas
	mux.Handle("/kelas", middleware.Auth(http.HandlerFunc(controller.KelasController)))
	mux.Handle("/tambah-kelas", middleware.Auth(http.HandlerFunc(controller.TambahKelasController)))
	mux.Handle("/store-kelas", middleware.Auth(http.HandlerFunc(controller.KelasStoreController)))
	mux.Handle("/edit-kelas", middleware.Auth(http.HandlerFunc(controller.KelasEditController)))
	mux.Handle("/update-kelas", middleware.Auth(http.HandlerFunc(controller.KelasUpdateController)))
	mux.Handle("/delete-kelas", middleware.Auth(http.HandlerFunc(controller.KelasDeleteController)))
	// End Kelas

	// start dosen
	mux.Handle("/dosen", middleware.Auth(http.HandlerFunc(controller.DosenController)))
	mux.Handle("/tambah-dosen", middleware.Auth(http.HandlerFunc(controller.DosenTambahController)))
	mux.Handle("/store-dosen", middleware.Auth(http.HandlerFunc(controller.DosenStoreController)))
	mux.Handle("/delete-dosen", middleware.Auth(http.HandlerFunc(controller.DosenDeleteController)))
	mux.Handle("/detail-dosen", middleware.Auth(http.HandlerFunc(controller.DosenDetailController)))
	mux.Handle("/edit-dosen", middleware.Auth(http.HandlerFunc(controller.DosenEditController)))
	mux.Handle("/update-dosen", middleware.Auth(http.HandlerFunc(controller.DosenUpdateController)))
	// end dosen

	// Kelas
	mux.Handle("/matkul", middleware.Auth(http.HandlerFunc(controller.MatkulController)))
	mux.Handle("/tambah-matkul", middleware.Auth(http.HandlerFunc(controller.TambahMatkulController)))
	mux.Handle("/store-matkul", middleware.Auth(http.HandlerFunc(controller.MatkulStoreController)))
	mux.Handle("/edit-matkul", middleware.Auth(http.HandlerFunc(controller.MatkulEditController)))
	mux.Handle("/update-matkul", middleware.Auth(http.HandlerFunc(controller.MatkulUpdateController)))
	mux.Handle("/delete-matkul", middleware.Auth(http.HandlerFunc(controller.MatkulDeleteController)))
	// End Kelas

	// Mahasiswa
	mux.Handle("/jadwal", middleware.Auth(http.HandlerFunc(controller.JadwalController)))
	mux.Handle("/tambah-jadwal", middleware.Auth(http.HandlerFunc(controller.TambahJadwalController)))
	mux.Handle("/store-jadwal", middleware.Auth(http.HandlerFunc(controller.StoreJadwalController)))
	// mux.Handle("/delete", middleware.Auth(http.HandlerFunc(controller.DeleteController)))
	// mux.Handle("/detail", middleware.Auth(http.HandlerFunc(controller.DetailController)))
	// mux.Handle("/edit", middleware.Auth(http.HandlerFunc(controller.EditController)))
	// mux.Handle("/update", middleware.Auth(http.HandlerFunc(controller.UpdateController)))
	// End Mahasiswa

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
