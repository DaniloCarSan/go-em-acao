// Usando append para acresscentar um elemento a fatia
package main

import "fmt"

func main() {

	slice := []int{10, 20, 30, 40, 50}

	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)

	newSlice := slice[1:3]

	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println(newSlice)

	newSlice = append(newSlice, 60)

	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println(newSlice)

	fmt.Println(slice)

}
