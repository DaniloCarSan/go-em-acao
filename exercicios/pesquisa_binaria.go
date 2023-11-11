package main

import "fmt"

func pesquisa(lista []int, search int) (int, error) {

	// 9
	//{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	//{6, 7, 8, 9, 10, 11}

	inicio := 0
	fim := len(lista) - 1

	for inicio <= fim {
		meio := inicio + fim/2

		chute := lista[meio]

		if chute == search {
			return meio, nil
		}

		if search > chute {
			inicio = meio + 1
		}

		if search < chute {
			fim = meio - 1
		}

	}

	return 0, fmt.Errorf("Nenhum item encontrado para busca: %d", search)
}

func main() {

	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	if index, err := pesquisa(lista, 2); err == nil {
		fmt.Printf("Numero encontrado na posição: %d \n", index)
	} else {
		fmt.Println(err.Error())
	}

}
