package main

import (
	"fmt"
	"time"
)

// ==========================================
// SELECT - Multiplexação de Canais em Go
// ==========================================
// Select é uma estrutura de controle que permite esperar
// múltiplas operações de comunicação simultâneas.
// É semelhante a um switch, mas para operações em canais.

// Exemplo 1: Select básico com múltiplos canais
func exemploSelect1() {
	fmt.Println("===== EXEMPLO 1: Select Básico =====")

	// Criando dois canais para comunicação
	canal1 := make(chan string)
	canal2 := make(chan string)

	// Goroutine que envia dados após 1 segundo
	go func() {
		time.Sleep(1 * time.Second)
		canal1 <- "Dados do canal 1"
	}()

	// Goroutine que envia dados após 2 segundos
	go func() {
		time.Sleep(2 * time.Second)
		canal2 <- "Dados do canal 2"
	}()

	// Select aguarda a primeira comunicação que está pronta
	// e executa o case correspondente
	select {
	case msg1 := <-canal1:
		fmt.Println("Recebido:", msg1)
	case msg2 := <-canal2:
		fmt.Println("Recebido:", msg2)
	}

	// Aguarda a segunda mensagem
	select {
	case msg1 := <-canal1:
		fmt.Println("Recebido:", msg1)
	case msg2 := <-canal2:
		fmt.Println("Recebido:", msg2)
	}
}

// Exemplo 2: Select com default (não bloqueante)
func exemploSelect2() {
	fmt.Println("\n===== EXEMPLO 2: Select com Default =====")

	canal := make(chan string)

	// Goroutine que envia dados após 3 segundos
	go func() {
		time.Sleep(3 * time.Second)
		canal <- "Finalmente chegou!"
	}()

	// O select tenta receber do canal
	// Se não houver dados prontos, executa o default (não bloqueia)
	select {
	case msg := <-canal:
		fmt.Println("Recebido:", msg)
	default:
		fmt.Println("Canal não está pronto, usando default!")
	}

	// Aguarda para receber a mensagem
	time.Sleep(4 * time.Second)

	select {
	case msg := <-canal:
		fmt.Println("Agora recebido:", msg)
	default:
		fmt.Println("Ainda não tem dados")
	}
}

// Exemplo 3: Select com timeout usando time.After
func exemploSelect3() {
	fmt.Println("\n===== EXEMPLO 3: Select com Timeout =====")

	canal := make(chan string)

	// Goroutine que envia dados após 5 segundos
	go func() {
		time.Sleep(5 * time.Second)
		canal <- "Chegou depois do timeout"
	}()

	// time.After retorna um canal que dispara após X tempo
	// Se o canal não responder em 2 segundos, executa o case timeout
	select {
	case msg := <-canal:
		fmt.Println("Recebido:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout! Nenhuma resposta em 2 segundos")
	}
}

// Exemplo 4: Select em loop monitorando múltiplas goroutines
func exemploSelect4() {
	fmt.Println("\n===== EXEMPLO 4: Select em Loop (Monitoramento) =====")

	// Canais para comunicação
	tarefasConcluidas := make(chan string)
	erros := make(chan string)
	parada := make(chan bool)

	// Goroutine 1: Simulando tarefas
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(1 * time.Second)
			tarefasConcluidas <- fmt.Sprintf("Tarefa %d concluída", i)
		}
	}()

	// Goroutine 2: Simulando possível erro após 4 segundos
	go func() {
		time.Sleep(4 * time.Second)
		erros <- "Erro detectado!"
	}()

	// Loop que monitora múltiplos canais
	tasksCompleted := 0
	for {
		select {
		case tarefa := <-tarefasConcluidas:
			fmt.Println("✓", tarefa)
			tasksCompleted++
			if tasksCompleted >= 3 {
				parada <- true
			}

		case erro := <-erros:
			fmt.Println("✗ Erro:", erro)

		case <-parada:
			fmt.Println("Todas as tarefas foram concluídas!")
			return
		}
	}
}

// Exemplo 5: Select com envio (sending) em canais
func exemploSelect5() {
	fmt.Println("\n===== EXEMPLO 5: Select com Envio em Canais =====")

	canal1 := make(chan string)
	canal2 := make(chan string)

	// Goroutine que recebe de ambos os canais
	go func() {
		for i := 0; i < 4; i++ {
			select {
			case msg := <-canal1:
				fmt.Println("Canal 1 recebeu:", msg)
			case msg := <-canal2:
				fmt.Println("Canal 2 recebeu:", msg)
			}
		}
	}()

	// Goroutine que alterna envios entre canais
	go func() {
		for i := 1; i <= 4; i++ {
			if i%2 == 0 {
				canal1 <- fmt.Sprintf("Mensagem %d do canal 1", i)
			} else {
				canal2 <- fmt.Sprintf("Mensagem %d do canal 2", i)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(3 * time.Second)
}

// Exemplo 6: Select com múltiplos casos e comportamento aleatório
func exemploSelect6() {
	fmt.Println("\n===== EXEMPLO 6: Select com Múltiplos Cases =====")

	mensagens := make(chan string)
	tick := time.Tick(500 * time.Millisecond)
	timeout := time.After(3 * time.Second)

	// Goroutine enviando mensagens
	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(800 * time.Millisecond)
			mensagens <- fmt.Sprintf("Mensagem #%d", i)
		}
	}()

	// Loop que trata múltiplos eventos
	for {
		select {
		case msg := <-mensagens:
			fmt.Println("📨 Nova mensagem:", msg)

		case <-tick:
			fmt.Println("⏱️ Tick a cada 500ms")

		case <-timeout:
			fmt.Println("⏰ Timeout atingido! Encerrando...")
			return
		}
	}
}

// Função principal que executa todos os exemplos
func main() {
	fmt.Println("========================================")
	fmt.Println("   EXEMPLOS DE SELECT EM GO")
	fmt.Println("========================================")

	exemploSelect1()
	exemploSelect2()
	exemploSelect3()
	exemploSelect4()
	exemploSelect5()
	exemploSelect6()

	fmt.Println("\n========================================")
	fmt.Println("   FIM DOS EXEMPLOS")
	fmt.Println("========================================")
}

// ==========================================
// RESUMO - PRINCIPAIS CARACTERÍSTICAS:
// ==========================================
// 1. SELECT - Multiplexação de canais
//    - Aguarda a primeira operação disponível
//    - Executa o case correspondente
//
// 2. CASES - Operações em canais
//    - Pode receber: case msg := <-canal
//    - Pode enviar: case canal <- msg
//
// 3. DEFAULT - Alternativa não-bloqueante
//    - Executado se nenhum case está pronto
//
// 4. TIMEOUT - Implementado com time.After
//    - Evita bloqueios indefinidos
//
// 5. LOOPS - Select em for
//    - Monitora múltiplos eventos simultaneamente
//
// 6. ORDEM - Quando múltiplos cases prontos
//    - A escolha é aleatória (fair random)
