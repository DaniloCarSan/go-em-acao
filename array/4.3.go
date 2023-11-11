// Declarando um array com GO calculando o tamanho automaticamente
package main

import "fmt"

func main() {

	elementos := [...]int{1, 2, 3, 4, 6}

	for _, value := range elementos {
		fmt.Println(value)
	}

}
