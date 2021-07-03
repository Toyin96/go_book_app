package main

import (
	"Go_Book_Git/loggers"
	"database/sql"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)


var tpl *template.Template
var database *sql.DB

func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
	db, err := sql.Open("mysql", "magna:m18job,,@tcp(127.0.0.1)/book")
	if err != nil{
		loggers.Error.Fatalln("couldn't establish a connection with the database", err)
	}
	defer db.Close()
}

func main(){
	router := mux.NewRouter()

	http.Handle("/", router)
	router.HandleFunc("/", home)
}