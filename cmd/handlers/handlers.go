package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/movieBlog/cmd/database"
)

type Film struct {
	Title string
	Director string
}

func HandlerHomepage(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", database.ConnectionString)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	
	movies, err := database.GetMovies(db)
	if err != nil {
		http.Error(w, "Error fetching movies", http.StatusInternalServerError)
		return
	}

	parsedHtml, err := template.ParseFiles("html/index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(parsedHtml, err)
	data := map[string][]database.Movie{
		"Movies": movies,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Template rendering error", http.StatusInternalServerError)
		return
	}
}