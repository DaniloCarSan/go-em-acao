// recuperando um valor de um mapa e testando dua ausência
package main

import "fmt"

func main() {
	notas := map[string]float64{}

	notas["Danilo"] = 8.5
	notas["Thais"] = 8
	notas["Rafael"] = 10

	// se a chave não exite retona zero para o tipo do valor do mapa que no cado é float = 0
	value := notas["joão"]

	fmt.Println(value)
}
