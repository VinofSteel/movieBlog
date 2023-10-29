package handlers

import (
	"html/template"
	"net/http"
)

type Film struct {
	Title string
	Director string
}

func HandlerHomepage(w http.ResponseWriter, r *http.Request) {
	parsedHtml, err := template.ParseFiles("html/index.html")
	
	tmpl := template.Must(parsedHtml, err)
	films := map[string][]Film{
		"Films": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "Blade Runner", Director: "Ridley Scott"},
			{Title: "The Thing", Director: "John Carpenter"},
		},
	}
	tmpl.Execute(w, films)
}