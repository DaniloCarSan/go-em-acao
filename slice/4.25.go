// Obtendo um fatia a partir de uma fatia
package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}

	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)

	// partir do elemento 2 pegue todos os elentos
	newSlice := slice[2:]

	fmt.Println(newSlice)

	//Quando um slice um slice é gerado a partir de outro slice
	// o novo slice continua apontando para os valores subjacentes do array que o slice antigo aponta
	// ou seja se eu alterar um valor de um determinado indice no novo slice
	// a alteração também será feita no slice original
	newSlice1 := slice[2:4]
	fmt.Println(newSlice1)

	newSlice1[0] = 233
	fmt.Println(newSlice1, slice)
}
