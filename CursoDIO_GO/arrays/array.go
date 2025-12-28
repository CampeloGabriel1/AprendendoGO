package main

import "fmt"

func main() {

	var x [5]float64

	x[0] = 7.5
	x[1] = 8.5
	x[2] = 6.0
	x[3] = 9.0
	x[4] = 5.5

	var total float64 = 0

	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	fmt.Println(total)
}