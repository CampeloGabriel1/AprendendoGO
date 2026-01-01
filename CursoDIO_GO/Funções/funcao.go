//função também é chamada de procediment ou sub-rotina

package main

import (
	"fmt"
)

func media (lista []float64) float64 {
	total := 0.0

	for _, valor := range lista {
		total += valor
	}
	return total/float64(len(lista))
}

func main() {
	/*lista := []float64{4, 5, 6, 7, 8}

	total := 0.0

	for _, valor := range lista {
		total += valor
	}
	fmt.Println(total/float64(len(lista)))*/

	lista := []float64{4, 5, 6, 7, 8}
	fmt.Println(media(lista))

				
}