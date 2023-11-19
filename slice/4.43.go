// compondo fatias de fatias
package main

import "fmt"

func main() {

	/// cria uma fatia de fatias de inteiros
	slice := [][]int{
		[]int{10},
		[]int{100, 200},
	}

	// adiciona valor a primeira fatia de inteiros

	slice[0] = append(slice[0], 300)

	fmt.Println(slice[0])
}
