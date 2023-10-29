package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/movieBlog/cmd/handlers"
)

func main() {
	port := ":8000"
	fmt.Printf("Server running in port %v\n", port)

	http.HandleFunc("/", handlers.HandlerHomepage)

	log.Fatal(http.ListenAndServe(port, nil))
}