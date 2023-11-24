//Declaração campos com base em outros tipos de estrutura
package main

import "fmt"

type User struct {
	name       string
	email      string
	ext        int
	priviliged bool
}

type Admin struct {
	person User
	level  string
}

func main() {

	gaya := Admin{
		person: User{
			name:       "Danilo",
			email:      "danilocarsan@gmail.com",
			ext:        10,
			priviliged: true,
		},
		level: "super",
	}

	fmt.Println(gaya)
}
