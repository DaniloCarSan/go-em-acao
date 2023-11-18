package main

import "fmt"

type Area struct {
	largura int
	altura  int
}

func (a Area) isBase() bool {
	return a.largura == a.altura
}

func (a Area) proximaArea() Area {

	if a.largura < a.altura {
		if a.altura%a.largura == 0 {
			a.altura = a.largura
			return a
		}
		return Area{largura: a.altura % a.largura, altura: a.largura}
	}

	if a.largura%a.altura == 0 {
		a.largura = a.altura
		return a
	}

	return Area{largura: a.largura % a.altura, altura: a.altura}
}

func calculaMaiorQuadrado(a Area) Area {

	if a.isBase() {
		return a
	}

	return calculaMaiorQuadrado(a.proximaArea())
}

func main() {

	fmt.Println(calculaMaiorQuadrado(Area{largura: 1680, altura: 640}))
	fmt.Println(calculaMaiorQuadrado(Area{largura: 900, altura: 90}))
	fmt.Println(calculaMaiorQuadrado(Area{largura: 81, altura: 9}))

}
