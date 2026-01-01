package main

import "fmt"

/*
CLOSURE EM GO

Closure é uma função que "captura" variáveis do escopo externo onde foi definida.
A função interna tem acesso às variáveis da função externa, mesmo após a função
externa ter terminado de executar. Isso permite que a função "lembre" do estado
dessas variáveis entre chamadas sucessivas.

Características principais:
1. A função interna "enxerga" as variáveis da função externa
2. Pode ler E modificar essas variáveis
3. Cada closure criado tem seu próprio estado independente
4. Muito útil para callbacks, decoradores, factory functions e data encapsulation

Exemplo real:
- Criar geradores de IDs (contador que incrementa)
- Implementar padrão Singleton
- Callbacks em servidores HTTP
- Função que retorna funções configuradas
*/

func main() {
	x:= 0

	numero := func() int {
		x++
		return x
	}

	fmt.Println(numero)
}