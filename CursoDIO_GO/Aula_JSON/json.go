//JSON = JavaScript Object Notation, formato padrão para enviar e receber dados em aplicações web. É leve, fácil de ler e escrever, e amplamente utilizado para comunicação entre servidores e clientes.//

package main

import (
	"encoding/json"
	"fmt"
	"io"
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
		fmt.Println("erro")

	}
	fmt.Println("Arquivo aberto com sucesso")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var usuarios Usuarios

	json.Unmarshal(byteValue, &usuarios)

	for i := 0; i < len(usuarios.Usuarios); i++ {
		fmt.Println(usuarios.Usuarios[i].Nome)
		fmt.Println(strconv.Itoa(usuarios.Usuarios[i].Idade))
		fmt.Println(usuarios.Usuarios[i].Email)
		fmt.Println(usuarios.Usuarios[i].Social)
	}

}