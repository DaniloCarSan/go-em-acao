//Declaração de uma variável do tipo  estrutura usando uma estrtura literal
package main

import "fmt"

type User struct {
	name       string
	email      string
	ext        int
	priviliged bool
}

func main() {

	bill := User{
		name:       "Danilo",
		email:      "danilocarsan@gmail.com",
		ext:        10,
		priviliged: true,
	}

	fmt.Println(bill)
}
