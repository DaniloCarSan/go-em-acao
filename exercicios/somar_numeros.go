package main

import "fmt"

func somar(numeros []int) int {

	if len(numeros) == 1 {
		return numeros[0]
	}

	return numeros[0] + somar(numeros[1:])
}

func main() {

	numeros := []int{1, 2, 12, 12, 12, 1, 12, 12, 12}

	fmt.Println(somar(numeros))

}
