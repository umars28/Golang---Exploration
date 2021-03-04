package config

import "net/http"

const (
	forbidden          = "Tidak diijinkan"
	internalServer     = "error is happening, keep calm"
	konfirmasiPassword = "Mohon maaf, Konfirmasi Password Harus Sama"
)

func MessageError503(w http.ResponseWriter, r *http.Request) {
	http.Error(w, forbidden, http.StatusNotFound)
	return
}

func MessageError500(w http.ResponseWriter, r *http.Request) {
	http.Error(w, internalServer, http.StatusInternalServerError)
	return
}

func ConfirmPassword(w http.ResponseWriter, r *http.Request) {
	http.Error(w, konfirmasiPassword, http.StatusInternalServerError)
	return
}
