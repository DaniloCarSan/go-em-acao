//Iterando pelo um mapa usando o for range
package main

import "fmt"

func main() {

	notas := map[string]float64{}

	notas["Danilo"] = 8.5
	notas["Thais"] = 8
	notas["Rafael"] = 10

	// o valor impresso são aleatorios
	// ou seja não uma ordem especifica de
	// impressão
	for k, v := range notas {
		fmt.Printf("%s =  %f \n", k, v)
	}
}
