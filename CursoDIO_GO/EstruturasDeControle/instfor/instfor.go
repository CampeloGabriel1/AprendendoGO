package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		/*fmt.Printf("O valor de i é %d\n", i)
			if i%2 == 0 {
		fmt.Println("O valor de i é par")
		} else {
		fmt.Println("O valor de i é ímpar")
		} */

		for j := 0; j <= 5; j++ {

			switch j {
			case 0: fmt.Println("O valor de i é zero")
			case 1: fmt.Println("O valor de i é um")
			case 2: fmt.Println("O valor de i é dois")
			case 3: fmt.Println("O valor de i é três")
			case 4: fmt.Println("O valor de i é quatro")
			case 5: fmt.Println("O valor de i é cinco")
			default:
			fmt.Printf("O valor de i não está entre zero e dez\n", i)
			}

		}

	}

}