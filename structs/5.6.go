//Declaração campos com base em outros tipos de estrutura
package main

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
