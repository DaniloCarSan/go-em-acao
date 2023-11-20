//Passando mapa entre funções
package main

import "fmt"

func main() {

	notas := map[string]float64{}

	notas["Danilo"] = 8.5
	notas["Thais"] = 8
	notas["Rafael"] = 10

	// ao passar um mapa para um função
	// o mesmo não é copiado
	// ele continua apontando para o mapa
	// externo assim como acontece par a slices
	deleteNota("Danilo", notas)

	for k, v := range notas {
		fmt.Printf("%s =  %f \n", k, v)
	}
}

func deleteNota(key string, notas map[string]float64) {
	delete(notas, key)
}
