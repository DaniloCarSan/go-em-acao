//Definindo uma variável do tipo struc definida como zero (para cada tipo)
package main

import "fmt"

type User struct {
	name       string
	email      string
	ext        int
	priviliged bool
}

func main() {

	var bill User

	fmt.Println(bill)
}
