package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/movieBlog/cmd/database"
	"github.com/movieBlog/cmd/handlers"
	"github.com/movieBlog/internal/environment"
)

func main() {
	database.InitializeDB()

	port := environment.EnvVariable("PORT")
	fmt.Printf("Server running in port %v\n", port)

	http.HandleFunc("/", handlers.HandlerHomepage)

	log.Fatal(http.ListenAndServe(port, nil))
}