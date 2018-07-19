package main

import (
	"net/http"
	"html/template"
	"github.com/julienschmidt/httprouter"
)

type HomePageData struct {
	PageTitle string
}

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl := template.Must(template.ParseFiles("views/home.gohtml"))
	
	data := HomePageData{
		PageTitle: "Welcome to the container",
	}
	
	tmpl.Execute(w, data)
}
