package main

import "fmt"

var a int = 42
var b string = "Hello, Go!"

func main() {

	var c float64 = float64(a)
	
	fmt.Println("Variável c é do tipo:", fmt.Sprintf("%T", c))
}
