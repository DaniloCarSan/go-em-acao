// Declarando slice de inteiro com tamanho e capacidade
package main

import "fmt"

func main() {

	// tamanho 3 capacidade 5
	slice := make([]int, 3, 5)

	for _, v := range slice {
		fmt.Println(v)
	}

}
