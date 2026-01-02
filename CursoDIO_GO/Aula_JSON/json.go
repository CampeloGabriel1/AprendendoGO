//JSON = JavaScript Object Notation, formato padrão para enviar e receber dados em aplicações web. É leve, fácil de ler e escrever, e amplamente utilizado para comunicação entre servidores e clientes.//

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Usuarios struct {
	Usuarios []Usuario `json:"usuarios "`
}

type Usuario struct {
	Nome     string `json:"nome"`
	Idade    int    `json:"idade"`
	Email    string `json:"email"`
	Social   Social `json:"sociais"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}
func main() {
	
	jsonFile, err := os.Open("usuarios.json")
	if err != nil {
		log.Fatalf("Erro ao abrir arquivo: %v\n", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Erro ao ler arquivo: %v\n", err)
	}

	var usuarios Usuarios

	err = json.Unmarshal(byteValue, &usuarios)
	if err != nil {
		log.Fatalf("Erro ao fazer parse do JSON: %v\n", err)
	}

	for i := 0; i < len(usuarios.Usuarios); i++ {
		fmt.Println(usuarios.Usuarios[i].Nome)
		fmt.Println(strconv.Itoa(usuarios.Usuarios[i].Idade))
		fmt.Println(usuarios.Usuarios[i].Email)
		fmt.Println(usuarios.Usuarios[i].Social)
	}

}