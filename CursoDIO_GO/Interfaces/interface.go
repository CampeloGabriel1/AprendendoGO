// CursoDIO_GO/Interfaces/interface.go
// Interfaces são coleções de métodos
//
// Em Go, uma interface define um conjunto de assinaturas de método.
// Um tipo implementa uma interface implicitamente ao definir esses métodos;
// não é necessário declarar explicitamente que o tipo implementa a interface.
// Interfaces permitem polimorfismo e desacoplamento: qualquer valor cujo tipo
// implemente os métodos exigidos pode ser usado onde a interface é esperada.
// A interface vazia `interface{}` aceita valores de qualquer tipo.

package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
	perimetro() float64
}

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}

func (q quadrado) perimetro() float64 {
	return 4 * q.lado
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func medir(g geometria) {
	fmt.Println(g)
	fmt.Println("Área:", g.area())
	fmt.Println("Perímetro:", g.perimetro())
}


func main() {

	q := quadrado{lado:4}
	c := circulo{raio:5}

	medir(q)
	medir(c)

}