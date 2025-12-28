//São Coleçoes de Campos

package main

import "fmt"

type Pessoa struct {
	nome string
	idade int

}

func main() {

	fmt.Println(Pessoa{"Ana", 54})
	fmt.Println(Pessoa{nome: "Carlos", idade: 30})
}