package main

import "fmt"
func main() {

	/*arr := [5]int {10,20,30,40,50}
	fatia:= arr[1:4]*/

	fatia1 := []int {1,2,3}
	fatia2 := append(fatia1, 4,5,6)
	fatia3 := make([]int, 2)
	copy(fatia3, fatia1)

	
	
	fmt.Println(fatia1)
	fmt.Println(fatia2)
	fmt.Println(fatia3)
}
