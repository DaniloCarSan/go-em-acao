// Erro de fiamento ao tentar fatiar uma capacidade mairo que a capacidade existe
package main

import "fmt"

func main() {

	frutas := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	newSlice := frutas[2:3:6]

	fmt.Println(newSlice)
}
