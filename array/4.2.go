// Declarando um array usando array literal
package main

import "fmt"

func main() {

	elementos := [3]int{1, 2, 3}

	for _, value := range elementos {
		fmt.Println(value)
	}

}
