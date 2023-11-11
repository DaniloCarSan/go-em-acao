// usando append para aumentar o tamanho e capacidade de uma fatia
package main

import "fmt"

func main() {

	slice := []int{10, 20, 30, 40}
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)

	fmt.Println("-----------------")
	// A capacidade dobra quando a capacidade é menor que mil
	// se for maior aumenta em 25% mas isso pode ser alterado
	//com o tempo
	newSlice := append(slice, 50)

	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println(newSlice)

	fmt.Println("-----------------")
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)

}
