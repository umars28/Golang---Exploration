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

func JadwalController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		jadwals, err := query.GetAllJadwal(ctx)

		if err != nil {
			fmt.Println(err)
		}
		template, err := template.ParseFiles(
			path.Join("views/jadwal", "jadwal.html"),
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

		err = template.Execute(w, jadwals)
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

func TambahJadwalController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		kelas, err := query.GetAllKelas(ctx)
		hari, err := query.GetAllHari(ctx)
		dosen, err := query.GetAllDosen(ctx)
		matkul, err := query.GetAllMatkul(ctx)
		var data = map[string]interface{}{
			"kelas":  kelas,
			"hari":   hari,
			"dosen":  dosen,
			"matkul": matkul,
		}

		// var Data []interface{}
		// Data = append(Data, kelas, hari, dosen, matkul)
		if err != nil {
			config.MessageError503(w, r)
			return
		}
		template, err := template.ParseFiles(
			path.Join("views/jadwal", "tambah.html"),
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

		err = template.Execute(w, data)
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

func StoreJadwalController(w http.ResponseWriter, r *http.Request) {}

// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
// 		return
// 	}

// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Println(err)
// 		config.MessageError500(w, r)
// 		return
// 	}

// 	nim := r.Form.Get("nim")
// 	name := r.Form.Get("name")
// 	semester := r.Form.Get("semester")
// 	email := r.Form.Get("email")
// 	kelas_id := r.Form.Get("kelas_id")
// 	query.CreateRow(name, nim, semester, email, kelas_id)
// 	fmt.Println("success")
// 	http.Redirect(w, r, "/mahasiswa", 302)
// 	w.Write([]byte("<script>alert('Sukses menambahkan data')</script>"))
// 	return
// }

// func DeleteController(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Query().Get("id")
// 	user_id := r.URL.Query().Get("userId")
// 	fmt.Println(id)
// 	fmt.Println(user_id)
// 	if id == "" {
// 		config.MessageError500(w, r)
// 		fmt.Println(w, "id tidak boleh kosong", http.StatusBadRequest)
// 		return
// 	}
// 	mhs, _ := strconv.Atoi(id)
// 	user, _ := strconv.Atoi(user_id)
// 	query.Delete(mhs, user)
// 	fmt.Println("sukses hapus data")
// 	http.Redirect(w, r, "/mahasiswa", 302)
// 	w.Write([]byte("<script>alert('Sukses menghapus data')</script>"))
// 	return
// }

// const (
// 	table = "mahasiswa"
// )

// func DetailController(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		id := r.URL.Query().Get("id")
// 		ctx, cancel := context.WithCancel(context.Background())
// 		defer cancel()
// 		mhs, _ := strconv.Atoi(id)
// 		mahasiswas, err := query.Detail(ctx, mhs)
// 		if err != nil {
// 			log.Println(err)
// 			config.MessageError500(w, r)
// 			return
// 		}
// 		template, err := template.ParseFiles(
// 			path.Join("views/mahasiswa", "detail.html"),
// 			path.Join("views/template", "main.html"),
// 			path.Join("views/template", "header.html"),
// 			path.Join("views/template", "sidebar.html"),
// 			path.Join("views/template", "footer.html"),
// 		)
// 		if err != nil {
// 			log.Println(err)
// 			config.MessageError500(w, r)
// 			return
// 		}

// 		err = template.Execute(w, mahasiswas)
// 		if err != nil {
// 			log.Println(err)
// 			config.MessageError500(w, r)
// 			return
// 		}
// 		return
// 	}

// 	config.MessageError503(w, r)
// 	return
// }

// func EditController(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		id := r.URL.Query().Get("id")
// 		mhs, _ := strconv.Atoi(id)
// 		ctx, cancel := context.WithCancel(context.Background())

// 		defer cancel()
// 		mahasiswas, err := query.Edit(ctx, mhs)
// 		template, err := template.ParseFiles(
// 			path.Join("views/mahasiswa", "edit.html"),
// 			path.Join("views/template", "main.html"),
// 			path.Join("views/template", "header.html"),
// 			path.Join("views/template", "sidebar.html"),
// 			path.Join("views/template", "footer.html"),
// 		)
// 		if err != nil {
// 			log.Println(err)
// 			config.MessageError500(w, r)
// 			return
// 		}

// 		err = template.Execute(w, mahasiswas)
// 		if err != nil {
// 			log.Println(err)
// 			config.MessageError500(w, r)
// 			return
// 		}
// 		return
// 	}

// 	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
// 	return
// }

// func UpdateController(w http.ResponseWriter, r *http.Request) {
// 	user_id := r.URL.Query().Get("userId")
// 	if r.Method != http.MethodPost {
// 		config.MessageError503(w, r)
// 		return
// 	}

// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Println(err)
// 		config.MessageError500(w, r)
// 		return
// 	}

// 	id := r.Form.Get("id")
// 	nim := r.Form.Get("nim")
// 	name := r.Form.Get("name")
// 	kelas := r.Form.Get("kelas")
// 	semester := r.Form.Get("semester")
// 	query.Update(id, nim, name, semester, kelas, user_id)
// 	fmt.Println("success")
// 	http.Redirect(w, r, "/mahasiswa", 302)
// 	w.Write([]byte("<script>alert('Sukses mengubah data')</script>"))
// 	return
// }
