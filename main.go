package main

import (
	"log"
	"net/http"

	"test/controllers/accountcontroller"
	"test/services"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server started on: http://localhost:1234")
	http.HandleFunc("/", services.Index)
	http.HandleFunc("/admin", services.Admin)
	http.HandleFunc("/show", services.Show)
	http.HandleFunc("/showu", services.Showu)
	http.HandleFunc("/new", services.New)
	http.HandleFunc("/edit", services.Edit)
	http.HandleFunc("/insert", services.Insert)
	http.HandleFunc("/update", services.Update)
	http.HandleFunc("/delete", services.Delete)
	http.HandleFunc("/login", accountcontroller.Login)
	http.HandleFunc("/salah", services.Salah)
	http.HandleFunc("/logout", accountcontroller.Logout)
	http.ListenAndServe(":1234", nil)
}
