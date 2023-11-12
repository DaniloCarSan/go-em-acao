// Calculando o tamanho do slice slic[2:3:4]
package main

import "fmt"

func main() {

	// i : inicio
	// j : fim (fim ocupado)
	// k : capacidade (fim real)
	//slice[i:j:k]

	// Tamanho = j - i = 3 -2 = 1
	// Capacidade = k - i = 4 - 2 = 2

	frutas := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	newSlice := frutas[2:3:4]

	fmt.Printf("Tamanho: %d \n", len(newSlice))
	fmt.Printf("Capacidade: %d \n", cap(newSlice))
	fmt.Println(newSlice)

}
