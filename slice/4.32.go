// declarando uam fatia de strings usando uma fatia literal
package main

import "fmt"

func main() {
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	fmt.Printf("Tamanho: %d \n", len(source))
	fmt.Printf("Capacidade: %d \n", cap(source))
}
