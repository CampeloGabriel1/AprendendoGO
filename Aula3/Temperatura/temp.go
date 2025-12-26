package main

import "fmt"

const ebulicaoF  = 212.0


func main() {

	tempF  := ebulicaoF

	tempC  := (tempF - 32) * 5 / 9

	fmt.Printf("A temperatura de ebulicão da água em F = %g e A temperatura de ebulicão da água em C = %g ", tempF, tempC)


}