//goroutines são funções ou métodos que são executados concorrentemente com outras goroutines dentro do mesmo espaço de endereço.

package main

import (
	"fmt"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go f(0)
	var escreva string
	fmt.Scanln(&escreva)
}