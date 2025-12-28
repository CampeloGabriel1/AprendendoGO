package main

import "fmt"

type Retangulo struct {
	altura  float64
	largura float64
}

func (r *Retangulo) area() float64 {
	return r.altura * r.largura
}

func main() {

	r := Retangulo{altura: 10.0, largura: 5.0}
	fmt.Println(r.area())
}