package services

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"test/config"
	"test/models"
	"text/template"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))
var tmpl = template.Must(template.ParseGlob("form/*"))

func Admin(w http.ResponseWriter, r *http.Request) {
	db := config.DbConn()
	sessions, _ := store.Get(r, "mysession")
	user := sessions.Values["user"]
	fmt.Println("user:", user)
	selDB, err := db.Query("SELECT * FROM katalog ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	katalog := models.Katalog{}
	res := []models.Katalog{}
	for selDB.Next() {
		var id int64
		var nama string
		var merk string
		var harga int64
		var model string
		var gambar string
		var size string
		err = selDB.Scan(&id, &nama, &merk, &harga, &model, &gambar, &size)
		if err != nil {
			panic(err.Error())
		}
		katalog.Id = id
		katalog.Nama = nama
		katalog.Merk = merk
		katalog.Harga = harga
		katalog.Model = model
		katalog.Gambar = gambar
		katalog.Size = size
		res = append(res, katalog)
	}
	tmpl.ExecuteTemplate(w, "Admin", res)
	defer db.Close()
}

func Index(w http.ResponseWriter, r *http.Request) {
	db := config.DbConn()
	selDB, err := db.Query("SELECT * FROM katalog ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	katalog := models.Katalog{}
	res := []models.Katalog{}
	for selDB.Next() {
		var id int64
		var nama string
		var merk string
		var harga int64
		var model string
		var gambar string
		var size string
		err = selDB.Scan(&id, &nama, &merk, &harga, &model, &gambar, &size)
		if err != nil {
			panic(err.Error())
		}
		katalog.Id = id
		katalog.Nama = nama
		katalog.Merk = merk
		katalog.Harga = harga
		katalog.Model = model
		katalog.Gambar = gambar
		katalog.Size = size
		res = append(res, katalog)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}
func Show(w http.ResponseWriter, r *http.Request) {
	db := config.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM katalog WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	katalog := models.Katalog{}
	for selDB.Next() {
		var id int64
		var nama string
		var merk string
		var harga int64
		var model string
		var gambar string
		var size string
		err = selDB.Scan(&id, &nama, &merk, &harga, &model, &gambar, &size)
		if err != nil {
			panic(err.Error())
		}
		katalog.Id = id
		katalog.Nama = nama
		katalog.Merk = merk
		katalog.Harga = harga
		katalog.Model = model
		katalog.Gambar = gambar
		katalog.Size = size
	}
	tmpl.ExecuteTemplate(w, "Show", katalog)
	defer db.Close()
}
func Showu(w http.ResponseWriter, r *http.Request) {
	db := config.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM katalog WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	katalog := models.Katalog{}
	for selDB.Next() {
		var id int64
		var nama string
		var merk string
		var harga int64
		var model string
		var gambar string
		var size string
		err = selDB.Scan(&id, &nama, &merk, &harga, &model, &gambar, &size)
		if err != nil {
			panic(err.Error())
		}
		katalog.Id = id
		katalog.Nama = nama
		katalog.Merk = merk
		katalog.Harga = harga
		katalog.Model = model
		katalog.Gambar = gambar
		katalog.Size = size
	}
	tmpl.ExecuteTemplate(w, "Showu", katalog)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Salah(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Salah", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := config.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM katalog WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	katalog := models.Katalog{}
	for selDB.Next() {
		var id int64
		var nama string
		var merk string
		var harga int64
		var model string
		var gambar string
		var size string
		err = selDB.Scan(&id, &nama, &merk, &harga, &model, &gambar, &size)
		if err != nil {
			panic(err.Error())
		}
		katalog.Id = id
		katalog.Nama = nama
		katalog.Merk = merk
		katalog.Harga = harga
		katalog.Model = model
		katalog.Gambar = gambar
		katalog.Size = size
	}
	tmpl.ExecuteTemplate(w, "Edit", katalog)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := config.DbConn()
	var gambar string
	if r.Method == "POST" {
		nama := r.FormValue("nama")
		merk := r.FormValue("merk")
		harga := r.FormValue("harga")
		model := r.FormValue("model")
		gambar = r.FormValue("gambar")
		size := r.FormValue("size")

		insForm, err := db.Prepare("INSERT INTO katalog(nama, merk, harga, model,gambar,size) VALUES(?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nama, merk, harga, model, gambar, size)
		log.Println("INSERT: Nama: " + nama + " | Merk: " + merk + " | Harga: " + harga + " | Model: " + model + " | Gambar: " + gambar + " | Size: " + size)
	}
	defer db.Close()
	http.Redirect(w, r, "admin", 301)

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := handler.Filename
	if gambar != "" {
		filename = fmt.Sprintf("%s%s", gambar, filepath.Ext(handler.Filename))
	}
	fileLocation := filepath.Join(dir, "/static/files", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func Update(w http.ResponseWriter, r *http.Request) {
	db := config.DbConn()
	if r.Method == "POST" {
		nama := r.FormValue("nama")
		merk := r.FormValue("merk")
		harga := r.FormValue("harga")
		model := r.FormValue("model")
		gambar := r.FormValue("gambar")
		size := r.FormValue("size")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE katalog SET nama=?, merk=?, harga=?,model=?,gambar=?, size=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nama, merk, harga, model, gambar, size, id)
		log.Println("UPDATE: Nama: " + nama + " | Merk: " + merk + " | Harga: " + harga + " | Model: " + model + " | Gambar: " + gambar + " | Size: " + size)
	}
	defer db.Close()
	http.Redirect(w, r, "admin", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := config.DbConn()
	katalog := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM katalog WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(katalog)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "admin", 301)
}
