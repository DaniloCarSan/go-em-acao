// Declarando uma fatia usando uma fatia literal
package main

import "fmt"

func main() {

	// O tamanho e capacidade são iguais ao numero de elementos inicializados
	slice := []string{"BLUE", "RED", "PINK", "ORANGE"}

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	for _, v := range slice {
		fmt.Println(v)
	}

}
