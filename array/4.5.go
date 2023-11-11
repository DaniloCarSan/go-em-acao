// Acessando elementos de um array
package main

import "fmt"

func main() {

	elementos := [5]int{2: 2, 4: 3}

	fmt.Println(elementos[2])
	fmt.Println(elementos[4])

	//altera um elemento especifico de um array de acordo com o indice
	elementos[2] = 33
	elementos[4] = 58

	fmt.Println(elementos[2])
	fmt.Println(elementos[4])
}
