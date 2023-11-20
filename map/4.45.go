// Declarando um mapa com make
package main

import "fmt"

func main() {

	// mapa vazio
	// um map com chave do tipi int e valor do tipo string
	dic := make(map[int]string)
	fmt.Println(dic)

	// cria um para com chave do tipo string e valor do tipo string
	// inializa o mapa com dois valores
	dict := map[string]string{"key1": "value1", "key2": "value2"}
	fmt.Println(dict)

}
