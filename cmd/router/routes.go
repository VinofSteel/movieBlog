package router

import (
	"net/http"

	"github.com/movieBlog/cmd/handlers"
)

func Router() {
	http.HandleFunc("/", handlers.HandlerHomepage)
	http.HandleFunc("/add-movie/", handlers.HandlerAddMovie)
}