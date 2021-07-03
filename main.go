package main

import (
	"Go_Book_Git/loggers"
	"Go_Book_Git/service"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

var database *sql.DB

func init(){
	db, err := sql.Open("mysql", "toyinboy:m18job,,@tcp(127.0.0.1:3306)/book")
	if err != nil{
		loggers.Error.Fatalln(err)
	}
	defer db.Close()
}

func main(){
	router := mux.NewRouter()

	http.Handle("/", router)
	router.HandleFunc("/", service.Home)
	loggers.Error.Fatalln(http.ListenAndServe(":8081", router))

}