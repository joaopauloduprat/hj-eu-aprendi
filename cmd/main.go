package main

import (
	"log"
	"net/http"
	"text/template"
)

type Contact struct {
	Name string
	Email string
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func newData() Data {
	return Data {
		Contacts: []Contact{
			newContact("John", "jd@gmail.com"),
			newContact("Clara", "cd@gmail.com"),
		},
	}
}

func newContact(name, email string) Contact {
	return Contact{
		Name: name,
		Email: email,
	}
}

var data = newData()

func handleCount (w http.ResponseWriter, r *http.Request) {

	block := "index"

	if (r.Method == "POST") {
		name := r.FormValue("name")
		email := r.FormValue("email")

		data.Contacts = append(data.Contacts, newContact(name,email))

		block = "display"
	}

	tmpl := template.Must(template.ParseFiles("../views/index.html"))

	tmpl.ExecuteTemplate(w, block, data)
}

func main() {

	http.HandleFunc("/", handleCount)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
