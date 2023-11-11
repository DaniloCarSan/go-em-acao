// Acessando elementos de um array de ponteiro
package main

import "fmt"

func main() {

	elementos := [5]*int{0: new(int), 1: new(int)}

	fmt.Println(elementos[0]) // imprime endereço da memoria
	fmt.Println(elementos[1]) // imprime endereço da memoria

	// Obtem os valores dos ponteiros  = 0 => tipo zero para numeros inteiros
	fmt.Println(*elementos[0])
	fmt.Println(*elementos[1])

	//Atribui valores
	*elementos[0] = 2
	*elementos[1] = 3

	fmt.Println(*elementos[0])
	fmt.Println(*elementos[1])
}
