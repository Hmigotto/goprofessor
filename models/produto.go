package models

import (
	"net/http"

	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Hmigotto/goprofessor/db"
	"go.mongodb.org/mongo-driver/bson"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      string
	Quantidade string
}

func Create(w http.ResponseWriter, r *http.Request) { //Esta linha define a função create com dois argumentos w e r do tipo http.ResponseWriter e *http.Request respectivamente. Esses dois tipos são usados para gerenciar a resposta da requisição HTTP.

	// Read the body of the POST request
	body, err := ioutil.ReadAll(r.Body)

	// Unmarshal the body into a Produto struct

	var produto Produto
	err = json.Unmarshal(body, &produto)
	if err != nil {
		log.Fatal(err)
	}

	// Create a person document
	_, err = db.ConectaComBancoDeDadosColectionPeople().InsertOne(context.TODO(), produto) //Esta linha insere o documento person na coleção people. O valor de retorno é armazenado na variável err.

	fmt.Println("Inserted person:", produto) //Esta linha imprime na resposta HTTP a mensagem "Inserted person" e os valores do documento person inserido.

}

func Display() []Produto {
	produtos := []Produto{}

	// Get all people from the database
	cursor, err := db.ConectaComBancoDeDadosColectionPeople().Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var produto Produto
		err := cursor.Decode(&produto)
		if err != nil {
			log.Fatal(err)
		}
		produtos = append(produtos, produto)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	cursor.Close(context.TODO())

	return produtos
}

func Remove(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")

		// Delete the document from the collection based on the "nome" field
		_, err := db.ConectaComBancoDeDadosColectionPeople().DeleteOne(context.TODO(), bson.M{"nome": nome})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Deleted document with nome:", nome)

		http.Redirect(w, r, "/", http.StatusSeeOther)

	}
}

func Updateproduto(w http.ResponseWriter, r *http.Request) { //Esta linha define a função create com dois argumentos w e r do tipo http.ResponseWriter e *http.Request respectivamente. Esses dois tipos são usados para gerenciar a resposta da requisição HTTP.

	// Read the body of the POST request
	body, err := ioutil.ReadAll(r.Body)

	// Unmarshal the body into a Produto struct

	var produto Produto
	err = json.Unmarshal(body, &produto)
	if err != nil {
		log.Fatal(err)
	}

	// Create a person document
	_, err = db.ConectaComBancoDeDadosColectionPeople().InsertOne(context.TODO(), produto) //Esta linha insere o documento person na coleção people. O valor de retorno é armazenado na variável err.

	fmt.Println("Updated person:", produto) //Esta linha imprime na resposta HTTP a mensagem "Inserted person" e os valores do documento person inserido.

}
