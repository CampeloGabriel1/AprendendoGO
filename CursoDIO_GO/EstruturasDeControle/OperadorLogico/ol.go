package main

import "fmt"

func main() {

	x:= 3
	if (x == 2 || x == 3) {
		fmt.Println("X é 2 ou 3")
	} else {
		fmt.Println("X não é 2 nem 3")
	}
}