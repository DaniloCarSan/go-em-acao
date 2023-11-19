//Passando fatias entre funções
package main

func main() {

	// sloca uma fatia com 1 milhão de inteiros
	slice := make([]int, 1e6)

	foo(slice)
}

func foo(slice []int) []int {

	return slice
}
