//Remover item do mapa com delete
package main

import "fmt"

func main() {

	notas := map[string]float64{}

	notas["Danilo"] = 8.5
	notas["Thais"] = 8
	notas["Rafael"] = 10

	delete(notas, "Danilo")

	for k, v := range notas {
		fmt.Printf("%s =  %f \n", k, v)
	}
}
