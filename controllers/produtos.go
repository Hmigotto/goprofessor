package controllers

import (
	"net/http"
	"text/template"

	"github.com/Hmigotto/goprofessor/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := models.Display()

	temp.ExecuteTemplate(w, "Index", produtos)
}

func Update(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Update", nil)
}
