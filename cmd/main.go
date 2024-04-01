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

func newContact(name, email string) Contact {
    return Contact{
        Name: name,
        Email: email,
    }
}

type Contacts = []Contact

type Data struct {
    Contacts Contacts
}

func (d *Data) hasEmail(email string) bool {
    for _, contact := range d.Contacts {
        if contact.Email == email {
            return true
        }
    }
    return false
}

func newData() Data {
    return Data{
        Contacts: []Contact{
            newContact("John", "aoeu"),
            newContact("Clara", "cd@gmail.com"),
        },
    }
}

type FormData struct {
    Values map[string]string
    Errors map[string]string
}

func newFormData() FormData {
    return FormData{
        Values: make(map[string]string),
        Errors: make(map[string]string),
    }
}

type Page struct {
    Data Data
    Form FormData
}

func newPage() Page {
    return Page{
        Data: newData(),
        Form: newFormData(),
    }
}


var page = newPage()

func handleContact(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("../views/index.html"))

	if (r.Method == "POST") {
		name := r.FormValue("name")
        email := r.FormValue("email")

        if page.Data.hasEmail(email) {
            formData := newFormData()
            formData.Values["name"] = name
            formData.Values["email"] = email
            formData.Errors["email"] = "Email already exists"

			w.WriteHeader(422)
			tmpl.ExecuteTemplate(w, "form", formData)
        } else {
			contact := newContact(name, email)
			page.Data.Contacts = append(page.Data.Contacts, contact)
	
			tmpl.ExecuteTemplate(w, "form", newFormData())
			tmpl.ExecuteTemplate(w, "oob-contact", contact)
		}
	} else {
		tmpl.ExecuteTemplate(w, "index", page)
	}
}

func main() {

	http.HandleFunc("/", handleContact)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
