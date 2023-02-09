package routes

import (
	"net/http"

	"github.com/Hmigotto/goprofessor/controllers"
	"github.com/Hmigotto/goprofessor/models"
)

func CarregaRotas() {

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/update", controllers.Update)

	http.HandleFunc("/create", models.Create)
	http.HandleFunc("/delete", models.Remove)
	http.HandleFunc("/updateproduto", models.Updateproduto)

}
