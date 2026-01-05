package main

import (
	"fmt"
)

func soma(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}

func multiplica(a int, b int) int {
	return a * b
}

func subtrai(a int, b int) int {
	return a - b
}

func divide(a int, b int) int {
	return a / b
}

func main() {
	x := soma(1,2,3)
	fmt.Println(x)
	y := multiplica(4,5)
	fmt.Println(y)
	z := subtrai(10,4)
	fmt.Println(z)
	w := divide(20,5)
	fmt.Println(w)

}

