// Como o tamanho e capacidade são calculados
// para slice recortados e outros array
package main

import "fmt"

func main() {

	sliceOriginal := []int{1, 2, 3, 4, 5}

	fmt.Println(sliceOriginal)
	fmt.Printf("Tamanho: %d \n", len(sliceOriginal))
	fmt.Printf("Tamanho: %d \n", cap(sliceOriginal))

	fmt.Println("------------SLICE-----RECORTADO")

	newSlice := sliceOriginal[1:3]

	fmt.Println(newSlice)
	fmt.Printf("Tamanho: %d \n", len(newSlice))
	fmt.Printf("Capacidade: %d \n", cap(newSlice))

	fmt.Printf("Tamanho: j(%d) - i(%d) = %d \n", 3, 1, len(newSlice))
	fmt.Printf("Capacidade: j(%d) - i(%d) = %d \n", 5, 1, cap(newSlice))
}
