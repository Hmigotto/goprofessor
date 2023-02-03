package main

import (
	"net/http"
	"text/template"

	"github.com/Hmigotto/goprofessor/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/create", models.Create)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := models.Display()

	temp.ExecuteTemplate(w, "Index", produtos)
}
