package controllers

import (
	//"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t := template.New("index")
	t.ParseFiles("templates/index.html")
	t.ExecuteTemplate(w, "index", nil)
}
