// Declarando um array inicializando elementos especificos
package main

import "fmt"

func main() {

	elementos := [5]int{2: 2, 4: 3}

	for _, value := range elementos {
		fmt.Println(value)
	}

}
