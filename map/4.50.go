// recuperando um valor de um mapa e testando dua ausência
package main

import "fmt"

func main() {
	notas := map[string]float64{}

	notas["Danilo"] = 8.5
	notas["Thais"] = 8
	notas["Rafael"] = 10

	if value, existe := notas["Danilo"]; existe {
		fmt.Println(value)
	}

	if value, existe := notas["joão"]; existe {
		fmt.Println(value)
	}
}
