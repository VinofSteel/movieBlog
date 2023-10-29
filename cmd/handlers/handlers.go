package handlers

import (
	"io"
	"net/http"
)

func HandlerHomepage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!\n")
	io.WriteString(w, r.Method)
}