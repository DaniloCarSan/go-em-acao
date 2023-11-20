//Declarando mapa que armazena fatias de string
package main

import "fmt"

func main() {

	dict := map[string][]string{"key1": []string{"GO", "PHP", "JAVA", "DART", "JAVASCRIPT"}}

	fmt.Println(dict["key1"])
}
