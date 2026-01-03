//API RESTful
//é um estilo arquitetural para sistemas distribuídos
//baseia-se em recursos (resources) e utiliza os verbos HTTP para realizar operações sobre esses recursos

//GET: listagem de registros
//POST: criação de registros
//DELETE: remoção de registros
//PUT: atualização de registros
//PATCH: atualização parcial de registros

//uma RESPONSE
//uma pokemon
//pokemonSpecies

package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/ditto")

	if err != nil {
		panic(err.Error())
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(responseData)) 



}

type response struct {

}

type pokemon struct {


}

type pokemonSpecies struct {
	Nome string `json:"name"`

}