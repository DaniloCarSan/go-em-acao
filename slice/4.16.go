// Cria uma fatia de string com o make
package main

import "fmt"

func main() {

	// slice com tamanho e capacidade igual a 5
	slice := make([]int, 5)

	for _, v := range slice {
		fmt.Println(v)
	}

}
