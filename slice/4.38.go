// Iterar pelas slices
package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for key, value := range s {
		fmt.Printf("key = %d | value= %d \n", key, value)
	}

}
