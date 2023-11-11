// Criando uma fatia com 3 indices
package main

import "fmt"

func main() {

	fmt.Println("----------------ORIGINAL---------------")
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	fmt.Printf("Tamanho: %d \n", len(source))
	fmt.Printf("Capacidade: %d \n", cap(source))
	fmt.Println(source)

	fmt.Println("----------------FATIA E RESTRINGE A CAPACIDADE---------------")

	slice := source[2:3:4]
	fmt.Printf("Tamanho: %d \n", len(slice))
	fmt.Printf("Capacidade: %d \n", cap(slice))
	fmt.Println(slice)

	slice[0] = "REPLACE"
	slice = append(slice, "APPEND")
	fmt.Println(source)
}
