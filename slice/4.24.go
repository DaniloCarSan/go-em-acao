// declarando uma fatia usanso uma fatia literal
package main

import "fmt"

func main() {

	// cria uma fatia de inteiros
	// Contem tamanho e capaciidade para 5 elementos

	slice := []int{10, 20, 30, 40, 50}

	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)

	// altera valor do elemento de indice 5
	slice[4] = 100

	fmt.Println(slice)
}
