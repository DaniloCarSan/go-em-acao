// Vantagems de definir o tamanho e a capacidade com o mesmo valor
package main

import "fmt"

func main() {

	// Tamanho e Capacidade = 5
	frutas := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	// Tamanho e Capacidade = 1
	slice := frutas[2:3:3]

	// Como o tamanho e a capacidade foram restringidas para o mesmo tamanho
	// quando o append é feito não há capacidade o suficiente isso faz com que
	// o slice de desvincule do array original gerando um novo exclusivo para este
	// slice
	slice = append(slice, "Uva")
	fmt.Println(frutas)
	fmt.Println(slice)

	slice[0] = "REPLACE"
	fmt.Println(frutas)
	fmt.Println(slice)

}
