package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Pessoa struct {
	ID 	 	int64	`json:"ID"`
	Nome      string	`json:"Nome"`
	Sobrenome string	`json:"Sobrenome"`
	Endereco  *Endereco `json:"Endereco"`
}

type Endereco struct {
	Cidade string `json:"Cidade"`
	Bairro string `json:"Bairro"`
	Rua string `json:"Rua"`
	Numero int64 `json:"Número"`
	CEP  string `json:"CEP"`
}

var pessoas []Pessoa

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")

}

func GetPeopleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	json.NewEncoder(w).Encode(pessoas)
}

func CreatePeopleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	var pessoa Pessoa

	err := decoder.Decode(&pessoa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validações simples
	if pessoa.Nome == "" || pessoa.Sobrenome == "" {
		http.Error(w, "Nome e Sobrenome são obrigatórios", http.StatusBadRequest)
		return
	}

	// Atribui um ID se não informado
	if pessoa.ID == 0 {
		pessoa.ID = time.Now().Unix()
	}

	pessoas = append(pessoas, pessoa)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pessoa)
}


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/clientes", GetPeopleHandler).Methods("GET")
	r.HandleFunc("/clientes", CreatePeopleHandler).Methods("POST")

	http.Handle("/", r)

	fmt.Println("Server listening on 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}