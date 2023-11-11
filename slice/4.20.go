// Declarando uma fatia com posição de indice
package main

import "fmt"

func main() {

	// Cria um slice com 100 posicoes a 99 é incializada com o valor 10
	// as demais são inicializadas com o valor zero do tipo que sendo int=0
	slice := []int{99: 10}

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	for i, v := range slice {
		fmt.Printf(" index= %d valor= %d \n", i, v)
	}

}
