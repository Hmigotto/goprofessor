package main

import (
	"net/http"

	"github.com/Hmigotto/goprofessor/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)

}
