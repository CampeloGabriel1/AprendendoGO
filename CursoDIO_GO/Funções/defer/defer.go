package main

import (
	"fmt"
)

// O comando `defer` agenda a execução de uma chamada de função para
// ocorrer quando a função que contém o `defer` terminar (retornar).
// Principais pontos:
// - É útil para liberar recursos (fechar arquivos, conexões, liberar mutexes).
// - As chamadas `defer` são executadas em ordem LIFO (last-in, first-out).
// - Mesmo se a função terminar por `return` ou por `panic`, os `defer`
//   ainda serão executados (antes do unwind do stack no caso de panic).

// Exemplo simples que demonstra a ordem LIFO dos `defer` e seu uso.
func main() {
	fmt.Println("Início da main")

	// Uso típico: agendar a liberação de um recurso.
	// Exemplo real (comentado):
	// f, _ := os.Open("arquivo.txt")
	// defer f.Close() // garante que o arquivo será fechado ao final

	// Múltiplos defers: serão executados na ordem inversa à que foram
	// declarados (LIFO).
	defer fmt.Println("defer 1: executa por último")
	defer fmt.Println("defer 2: executa primeiro entre os defers")

	fmt.Println("Antes de sair da main (os defers ainda não executaram)")
	fmt.Println("Fim da main — agora retornamos e os defers executam")
}

