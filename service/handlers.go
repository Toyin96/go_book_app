package service

import (
	"Go_Book_Git/loggers"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func Home(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w, "home.gohtml", nil)
	if err != nil{
		loggers.Error.Fatalln(err)
	}
}

func CreateBook(w http.ResponseWriter, req *http.Request){

}