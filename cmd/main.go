package main

import (
	"net/http"
	"text/template"
)

type Templates struct {
	templates *template.Template
}

func main() {
	http.ListenAndServe(":8080", nil)
}