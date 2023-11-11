// Usa o make para criar uma fatia vazia de inteiros vazia
package main

import "fmt"

func main() {
	slice := make([]int, 0)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)
}
