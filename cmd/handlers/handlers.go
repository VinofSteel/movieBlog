package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/movieBlog/cmd/database"
)

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

func HandlerAddMovie(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", database.ConnectionString)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	title, director := r.PostFormValue("title"), r.PostFormValue("director")
	movieStruct := database.Movie{Title: title, Director: director}

	id := database.InsertMovie(db, movieStruct)
	fmt.Printf("Movie %v added to the DB with the id %v\n", movieStruct.Title, id)

	parsedHtml, err := template.ParseFiles("html/index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(parsedHtml, err)
	tmpl.ExecuteTemplate(w, "movie-list-element", movieStruct)
}