// Usando identificador vazio para ignorar o valor do índice
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 29, 10, 11, 12, 213, 14, 152, 16, 17, 18, 129, 20}

	for _, value := range s {
		fmt.Printf("value= %d \n", value)
	}
}
