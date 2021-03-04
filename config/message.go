package config

import "net/http"

const (
	forbidden = "Tidak diijinkan"
)

func MessageError503(w http.ResponseWriter, r *http.Request) {
	http.Error(w, forbidden, http.StatusNotFound)
	return
}
